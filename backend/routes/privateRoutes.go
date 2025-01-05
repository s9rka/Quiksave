package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	auth "notas/authentication"
	"notas/database"
	"notas/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

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

    w.WriteHeader(http.StatusNoContent) // 204 No Content
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

    // Fetch only the necessary user details (username, email)
    user, err := database.GetUserByID(userID)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
        }
        return
    }

    // Respond with user details
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
        return
    }
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

    // Update the note
    err = database.UpdateNoteInDB(noteID, userID, noteData)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            http.Error(w, "Note not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    // Fetch the updated note (including last_edit)
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

	// Extract userID from the context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Fetch tags for the user
	tags, err := database.GetNoteTags(userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Tags not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch tags", http.StatusInternalServerError)
		}
		return
	}

	// Respond with tags as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		http.Error(w, "Failed to encode tags data", http.StatusInternalServerError)
		return
	}
}