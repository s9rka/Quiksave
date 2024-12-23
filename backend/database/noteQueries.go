package database

import (
	"context"
	"fmt"
	"notas/models"
	"strings"
	"time"
)

func AddNewNoteToDB(note models.Note) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var noteID int
	newNoteQuery := `INSERT INTO notes (user_id, title, content) VALUES ($1, $2, $3) RETURNING id`

	err := dbPool.QueryRow(ctx, newNoteQuery, note.UserID, note.Title, note.Content).Scan(&noteID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert note: %w", err)
	}

	tagIDs := make([]int, 0)
	for _, tag := range note.Tags {
		var tagID int
		tagQuery := `
			INSERT INTO tags (name, user_id) VALUES ($1, $2)
			ON CONFLICT (name, user_id) DO NOTHING
			RETURNING id`
		err := dbPool.QueryRow(ctx, tagQuery, tag, note.UserID).Scan(&tagID)
		if err != nil {
			// If no ID was returned (tag already exists), fetch the ID
			if strings.Contains(err.Error(), "no rows in result set") {
				err = dbPool.QueryRow(ctx, `SELECT id FROM tags WHERE name = $1 AND user_id = $2`, tag, note.UserID).Scan(&tagID)
				if err != nil {
					return 0, fmt.Errorf("failed to get tag ID for '%s': %w", tag, err)
				}
			} else {
				return 0, fmt.Errorf("failed to insert tag '%s': %w", tag, err)
			}
		}
		tagIDs = append(tagIDs, tagID)
	}

	// Insert associations into  'note_tags' table
	for _, tagID := range tagIDs {
		noteTagQuery := `INSERT INTO note_tags (note_id, tag_id) VALUES ($1, $2)`
		_, err := dbPool.Exec(ctx, noteTagQuery, noteID, tagID)
		if err != nil {
			return 0, fmt.Errorf("failed to associate note with tag ID %d: %w", tagID, err)
		}
	}

	return noteID, nil
}

func GetNotesFromDB(userID int) ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	getNotesQuery := `SELECT n.id, n.title, n.content, n.created_at, COALESCE(ARRAY_AGG(COALESCE(t.name, '')), '{}') AS tags 
                        FROM notes n
						LEFT JOIN note_tags nt ON n.id = nt.note_id
						LEFT JOIN tags t ON nt.tag_id = t.id
						WHERE n.user_id = $1
						GROUP BY n.id
						ORDER BY n.created_at DESC
						`
	rows, err := dbPool.Query(ctx, getNotesQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notes: %w", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		var tags []string

		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &tags)
		if err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}

		if tags == nil {
			tags = []string{}
		}

		note.Tags = tags
		note.UserID = userID
		notes = append(notes, note)
	}

	return notes, nil
}

func GetNoteByIDFromDB(noteID, userID int) (*models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
		SELECT n.id, n.title, n.content, n.created_at, COALESCE(ARRAY_AGG(t.name), '{}') AS tags
		FROM notes n
		LEFT JOIN note_tags nt ON n.id = nt.note_id
		LEFT JOIN tags t ON nt.tag_id = t.id
		WHERE n.id = $1 AND n.user_id = $2
		GROUP BY n.id
	`

	var note models.Note
	var tags []string

	err := dbPool.QueryRow(ctx, query, noteID, userID).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &tags)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, fmt.Errorf("note with ID %d not found", noteID)
		}
		return nil, fmt.Errorf("failed to fetch note by ID %d: %w", noteID, err)
	}

	note.Tags = tags
	note.UserID = userID
	return &note, nil
}

func DeleteNoteFromDB(noteID, userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
        DELETE FROM notes
        WHERE id = $1 AND user_id = $2
        RETURNING id
    `

	var deletedNoteID int
	err := dbPool.QueryRow(ctx, query, noteID, userID).Scan(&deletedNoteID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("note with ID %d not found", noteID)
		}
		return fmt.Errorf("failed to delete note with ID %d: %w", noteID, err)
	}

	return nil
}
