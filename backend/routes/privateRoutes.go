package routes

import (
	"encoding/json"
	"net/http"
	"notas/database"
	"notas/models"
)

func NewNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	noteID, err := database.CreateNote(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"noteID": noteID,
		"message": "Note created successfully!",
	}

	json.NewEncoder(w).Encode(response)

}
