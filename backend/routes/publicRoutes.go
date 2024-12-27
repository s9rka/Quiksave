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
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(200)
}

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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var loginRequest LoginRequest
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

    // Refresh token
    refreshTokenValue := fmt.Sprintf("userID:%d|type:refresh", userID)
    refreshCookie := auth.CreateCookie("refresh_token", refreshTokenValue, "/", 7*24*time.Hour)
    if err := auth.WriteEncrypted(w, refreshCookie, []byte(secretKey)); err != nil {
        log.Printf("Failed to set encrypted refresh token cookie: %v", err)
        http.Error(w, "Failed to set encrypted refresh token cookie", http.StatusInternalServerError)
        return
    }

    // Access token
    accessTokenValue := fmt.Sprintf("userID:%d|type:access", userID)
    accessCookie := auth.CreateCookie("access_token", accessTokenValue, "/", 15*time.Minute)
    if err := auth.WriteEncrypted(w, accessCookie, []byte(secretKey)); err != nil {
        log.Printf("Failed to set encrypted access token cookie: %v", err)
        http.Error(w, "Failed to set encrypted access token cookie", http.StatusInternalServerError)
        return
    }

    // Log for debugging
    log.Printf("Access Token Cookie: %s\nRefresh Token Cookie: %s", accessCookie.Value, refreshCookie.Value)

    // Response
    response := LoginResponse{
        Message: "Login successful!",
    }
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Printf("Failed to encode response: %v", err)
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

