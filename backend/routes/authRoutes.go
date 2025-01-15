package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	auth "notas/authentication"
	"notas/database"
	"notas/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Missing required fields: username, email, or password.", http.StatusBadRequest)
		return
	}

	userID, err := database.AddUserToDB(user)
	if err != nil {
		if errors.Is(err, database.ErrDuplicateEmail) {
			http.Error(w, "Email already exists. Please use a different email.", http.StatusConflict)
			return
		}
		if errors.Is(err, database.ErrDuplicateUsername) {
			http.Error(w, "Username already exists. Please choose a different username.", http.StatusConflict)
			return
		}
		log.Printf("Failed to add user to database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"userID":  userID,
		"message": "User successfully created!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var loginRequest models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := database.ValidateUserLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, "Wrong username or password", http.StatusUnauthorized)
		return
	}

	secretKey := os.Getenv("SECRET_KEY")

	refreshTokenValue := fmt.Sprintf("userID:%d|type:refresh", userID)
	refreshCookie := auth.CreateCookie("refresh_token", refreshTokenValue, "/", 7*24*time.Hour)
	if err := auth.WriteEncrypted(w, refreshCookie, []byte(secretKey)); err != nil {
		log.Printf("Failed to set encrypted refresh token cookie: %v", err)
		http.Error(w, "Failed to set encrypted refresh token cookie", http.StatusInternalServerError)
		return
	}

	accessTokenValue := fmt.Sprintf("userID:%d|type:access", userID)
	accessCookie := auth.CreateCookie("access_token", accessTokenValue, "/", 15*time.Minute)
	if err := auth.WriteEncrypted(w, accessCookie, []byte(secretKey)); err != nil {
		log.Printf("Failed to set encrypted access token cookie: %v", err)
		http.Error(w, "Failed to set encrypted access token cookie", http.StatusInternalServerError)
		return
	}

	response := models.LoginResponse{
		Message: "Login successful!",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),
        Path:     "/",
        HttpOnly: true,
        Secure:   false,
        SameSite: http.SameSiteNoneMode,
    })

	http.SetCookie(w, &http.Cookie{
        Name:     "access_token",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour),
        Path:     "/",
        HttpOnly: true,
        Secure:   false,
        SameSite: http.SameSiteNoneMode,
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"Logged out successfully!"}`))
}


// Checks if there's avalid refresh token in cookie
// Gathers the userID from the refresh token
// Generates new JWT token with userID
func RefreshJWT(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh token not found", http.StatusUnauthorized)
		return
	}

	refreshToken := cookie.Value

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}
	userIDStr := claims["sub"].(string)

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	newJWT, err := auth.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"accessToken":"` + newJWT + `"}`))
}

func GetMe(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    user, err := database.GetUserByID(userID)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
        return
    }
}