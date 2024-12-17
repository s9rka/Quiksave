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

	response := map[string]interface{} {
		"userID": userID,
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
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
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

	accessToken, err := auth.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := auth.GenerateRefreshToken(userID)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	response := LoginResponse{
		Message: "Login successful!",
		AccessToken: accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}