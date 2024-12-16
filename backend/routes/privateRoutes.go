package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notas/database"
	"notas/models"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	noteID, err := database.AddNewNoteToDB(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"noteID": noteID,
		"message": "Note created successfully!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}

}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	notes, err := database.GetNotesFromDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get notes: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode notes: %v", err), http.StatusInternalServerError)
	}
}
