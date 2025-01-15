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
	newNoteQuery := `INSERT INTO notes (user_id, heading, content) VALUES ($1, $2, $3) RETURNING id`

	err := dbPool.QueryRow(ctx, newNoteQuery, note.UserID, note.Heading, note.Content).Scan(&noteID)
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

	getNotesQuery := `SELECT n.id, n.heading, n.content, n.created_at, n.last_edit, COALESCE(ARRAY_AGG(COALESCE(t.name, '')), '{}') AS tags 
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

		err := rows.Scan(&note.ID, &note.Heading, &note.Content, &note.CreatedAt, &note.LastEdit, &tags)
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
		SELECT n.id, n.heading, n.content, n.created_at, n.last_edit, COALESCE(ARRAY_AGG(COALESCE(t.name, '')), '{}') AS tags
		FROM notes n
		LEFT JOIN note_tags nt ON n.id = nt.note_id
		LEFT JOIN tags t ON nt.tag_id = t.id
		WHERE n.id = $1 AND n.user_id = $2
		GROUP BY n.id
	`

	var note models.Note
	var tags []string

	err := dbPool.QueryRow(ctx, query, noteID, userID).Scan(&note.ID, &note.Heading, &note.Content, &note.CreatedAt, &note.LastEdit, &tags)
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

func UpdateNoteInDB(noteID, userID int, noteData models.Note) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    tx, err := dbPool.Begin(ctx)
    if err != nil {
        return fmt.Errorf("error starting transaction: %w", err)
    }
    defer func() {
        if err != nil {
            _ = tx.Rollback(ctx)
        } else {
            err = tx.Commit(ctx)
        }
    }()

    updateNoteQuery := `
        UPDATE notes
        SET heading = $1,
            content = $2,
            last_edit = CURRENT_TIMESTAMP
        WHERE id = $3 AND user_id = $4
        RETURNING id
    `
    var updatedID int
    err = tx.QueryRow(ctx, updateNoteQuery, noteData.Heading, noteData.Content, noteID, userID).Scan(&updatedID)
    if err != nil {
        if strings.Contains(err.Error(), "no rows in result set") {
            return fmt.Errorf("note with ID %d not found", noteID)
        }
        return fmt.Errorf("failed to update note (ID %d): %w", noteID, err)
    }

    // Remove old tag associations
    deleteTagsQuery := `DELETE FROM note_tags WHERE note_id = $1`
    _, err = tx.Exec(ctx, deleteTagsQuery, noteID)
    if err != nil {
        return fmt.Errorf("failed to delete old tag associations: %w", err)
    }

    // Upsert new tags
    tagIDs := make([]int, 0, len(noteData.Tags))
    for _, tagName := range noteData.Tags {
        var tagID int
        insertOrIgnoreTagQuery := `
            INSERT INTO tags (name, user_id)
            VALUES ($1, $2)
            ON CONFLICT (name, user_id) DO NOTHING
            RETURNING id
        `
        err = tx.QueryRow(ctx, insertOrIgnoreTagQuery, tagName, userID).Scan(&tagID)
        if err != nil {
            // If tag already exists, fetch its ID
            if strings.Contains(err.Error(), "no rows in result set") {
                getExistingTagIDQuery := `
                    SELECT id FROM tags WHERE name = $1 AND user_id = $2
                `
                err = tx.QueryRow(ctx, getExistingTagIDQuery, tagName, userID).Scan(&tagID)
                if err != nil {
                    return fmt.Errorf("failed to get existing tag ID for '%s': %w", tagName, err)
                }
            } else {
                return fmt.Errorf("failed to upsert tag '%s': %w", tagName, err)
            }
        }
        tagIDs = append(tagIDs, tagID)
    }

    // Re-associate tags in note_tags
    insertNoteTagQuery := `INSERT INTO note_tags (note_id, tag_id) VALUES ($1, $2)`
    for _, tid := range tagIDs {
        _, err = tx.Exec(ctx, insertNoteTagQuery, noteID, tid)
        if err != nil {
            return fmt.Errorf("failed to associate note with tag ID %d: %w", tid, err)
        }
    }

	deleteUnusedTagsQuery := `DELETE FROM tags WHERE id NOT IN ( SELECT DISTINCT tag_id FROM note_tags) AND user_id = $1`
    _, err = tx.Exec(ctx, deleteUnusedTagsQuery, userID)
    if err != nil {
        return fmt.Errorf("failed to delete unused tags: %w", err)
    }

    return nil
}

func GetNoteTags(userID int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
		SELECT DISTINCT t.name
		FROM tags t
		WHERE t.user_id = $1
		ORDER BY t.name
	`

	rows, err := dbPool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tags for user ID %d: %w", userID, err)
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("failed to scan tag: %w", err)
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return tags, nil
}