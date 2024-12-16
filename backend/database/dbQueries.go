package database

import (
	"context"
	"fmt"
	"notas/models"
	"strings"
	"time"

)

func CreateNote(note models.Note) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	var noteID int
	newNoteQuery := "INSERT INTO notes (title, content, tags) VALUES ($1, $2, $3) RETURNING id"

	err := dbPool.QueryRow(ctx, newNoteQuery, note.Title, note.Content, note.Tags).Scan(&noteID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert note: %w", err)
	}

	tagIDs := make([]int, 0)
	for _, tag := range note.Tags {
		var tagID int
		tagQuery := `
			INSERT INTO tags (name) VALUES ($1)
			ON CONFLICT (name) DO NOTHING
			RETURNING id`
		err := dbPool.QueryRow(ctx, tagQuery, tag).Scan(&tagID)
		if err != nil {
			// If no ID was returned (tag already exists), fetch the ID
			if strings.Contains(err.Error(), "no rows in result set") {
				err = dbPool.QueryRow(ctx, "SELECT id FROM tags WHERE name = $1", tag).Scan(&tagID)
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
		noteTagQuery := "INSERT INTO note_tags (note_id, tag_id) VALUES ($1, $2)"
		_, err := dbPool.Exec(ctx, noteTagQuery, noteID, tagID)
		if err != nil {
			return 0, fmt.Errorf("failed to associate note with tag ID %d: %w", tagID, err)
		}
	}

	return noteID, nil
}