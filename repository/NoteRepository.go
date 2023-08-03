package repository

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/google/uuid"
)

// InsertNote inserts a new note to the database
func (repo *Repo) InsertNote(userId string, title string, content string) error {
	// Generate uuid for note_id
	noteId := uuid.New().String()

	stmt, err := repo.db.Prepare("INSERT INTO notes (note_id, user_id, title, content) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	_, err = stmt.Exec(noteId, userId, title, content)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	fmt.Println("note inserted to database")
	return nil
}

// GetNotes returns all notes of a user
func (repo *Repo) GetNotes(userId string) ([]models.Note, error) {
	var noteList []models.Note

	rows, err := repo.db.Query("SELECT note_id, title, content, created_at FROM notes WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		noteList = append(noteList, note)
	}

	return noteList, nil
}

// GetNoteByID returns a note by its id
func (repo *Repo) GetNoteByID(sessId, noteId string) (*models.Note, error) {
	var note models.Note

	// Check if note belongs to user
	stmt, err := repo.db.Prepare("SELECT note_id, title, content, created_at FROM notes WHERE note_id = $1 AND user_id = $2")
	if err != nil {
		return nil, fmt.Errorf("statement prepare error: %w", err)
	}

	err = stmt.QueryRow(noteId, sessId).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("query row error: %w", err)
	}

	return &note, nil
}

// UpdateNoteByID updates a note by its id
func (repo *Repo) UpdateNoteByID(sessId, noteId, title, content string) error {
	// Check if note belongs to user
	stmt, err := repo.db.Prepare("UPDATE notes SET title = $1, content = $2 WHERE note_id = $3 AND user_id = $4")
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	_, err = stmt.Exec(title, content, noteId, sessId)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	return nil
}

// DeleteNoteByID deletes a note by its id
func (repo *Repo) DeleteNoteByID(sessId, noteId string) error {
	// Check if note belongs to user
	stmt, err := repo.db.Prepare("DELETE FROM notes WHERE note_id = $1 AND user_id = $2")
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	_, err = stmt.Exec(noteId, sessId)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	return nil
}
