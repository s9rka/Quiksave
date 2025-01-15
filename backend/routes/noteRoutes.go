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


func CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	note.UserID = userID

	noteID, err := database.AddNewNoteToDB(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"noteID":  noteID,
		"message": "Note created successfully!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	notes, err := database.GetNotesFromDB(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get notes: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode notes: %v", err), http.StatusInternalServerError)
		return
	}
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
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
	noteIDStr := vars["id"]
	if noteIDStr == "" {
		http.Error(w, "Missing note ID", http.StatusBadRequest)
		return
	}

	noteID, err := strconv.Atoi(noteIDStr)
	if err != nil {
		http.Error(w, "Invalid note ID format", http.StatusBadRequest)
		return
	}

	note, err := database.GetNoteByIDFromDB(noteID, userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Note not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch note", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    vars := mux.Vars(r)
    noteIDStr := vars["id"]
    if noteIDStr == "" {
        http.Error(w, "Missing note ID", http.StatusBadRequest)
        return
    }

    noteID, err := strconv.Atoi(noteIDStr)
    if err != nil {
        http.Error(w, "Invalid note ID format", http.StatusBadRequest)
        return
    }

    err = database.DeleteNoteFromDB(noteID, userID)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            http.Error(w, "Note not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to delete note", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusNoContent)
}


func EditNote(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    vars := mux.Vars(r)
    noteIDStr := vars["id"]
    if noteIDStr == "" {
        http.Error(w, "Missing note ID", http.StatusBadRequest)
        return
    }

    noteID, err := strconv.Atoi(noteIDStr)
    if err != nil {
        http.Error(w, "Invalid note ID format", http.StatusBadRequest)
        return
    }

    var noteData models.Note
    if err := json.NewDecoder(r.Body).Decode(&noteData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    
    err = database.UpdateNoteInDB(noteID, userID, noteData)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            http.Error(w, "Note not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    updatedNote, err := database.GetNoteByIDFromDB(noteID, userID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to retrieve updated note: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(updatedNote)
}


func GetUserTags(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tags, err := database.GetNoteTags(userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Tags not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch tags", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "Failed to encode tags data", http.StatusInternalServerError)
		return
	}
}