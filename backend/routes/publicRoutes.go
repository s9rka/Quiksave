package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"notas/database"
	"notas/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var loginRequest LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := database.ValidateUserLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, "Wrong username or password", http.StatusUnauthorized)
	}

	response := fmt.Sprintf("User %d successfully logged in!", userID)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}