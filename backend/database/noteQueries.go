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
	newNoteQuery := `INSERT INTO notes (title, content, tags) VALUES ($1, $2, $3) RETURNING id`

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
				err = dbPool.QueryRow(ctx, `SELECT id FROM tags WHERE name = $1`, tag).Scan(&tagID)
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

func GetNotesFromDB() ([]models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	getNotesQuery := `SELECT n.id, n.title, n.content, n.created_at, COALESCE(ARRAY_AGG(t.name), '{}') AS tags FROM notes n
						LEFT JOIN note_tags nt ON n.id = nt.note_id
						LEFT JOIN tags t ON nt.tag_id = t.id
						GROUP BY n.id
						ORDER BY n.created_at DESC
						`
	rows, err := dbPool.Query(ctx, getNotesQuery)
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

		note.Tags = tags
		notes = append(notes, note)
	}

	return notes, nil
}

