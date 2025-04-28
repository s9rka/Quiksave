package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notas/database"
	"notas/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateVault(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var vault models.Vault
	if err := json.NewDecoder(r.Body).Decode(&vault); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vault.UserID = userID

	vaultID, err := database.CreateVault(vault)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create vault: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"vaultID": vaultID,
		"message": "Vault created successfully!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

func GetVaults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vaults, err := database.GetVaults(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get vaults: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vaults); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode vaults: %v", err), http.StatusInternalServerError)
		return
	}
}

func GetVaultByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	vaultIDStr := vars["id"]
	if vaultIDStr == "" {
		http.Error(w, "Missing vault ID", http.StatusBadRequest)
		return
	}

	vaultID, err := strconv.Atoi(vaultIDStr)
	if err != nil {
		http.Error(w, "Invalid vault ID format", http.StatusBadRequest)
		return
	}

	vault, err := database.GetVaultByID(vaultID, userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Vault not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Failed to fetch vault: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vault); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode vault: %v", err), http.StatusInternalServerError)
		return
	}
}
