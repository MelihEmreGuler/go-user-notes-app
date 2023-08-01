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

	stmt, err := repo.db.Prepare("insert into notes (note_id, user_id, title, content) values ($1, $2, $3, $4)")
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

// List returns all notes of a user
func (repo *Repo) List(userId string) ([]models.Note, error) {
	var noteList []models.Note

	rows, err := repo.db.Query("select note_id, title, content, created_at from notes where user_id = $1", userId)
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

// SelectById returns a note by its id
func (repo *Repo) SelectById(noteId string) (*models.Note, error) {
	var note models.Note

	stmt, err := repo.db.Prepare("select note_id, title, content, created_at from notes where note_id = $1")
	if err != nil {
		return nil, fmt.Errorf("statement prepare error: %w", err)
	}

	err = stmt.QueryRow(noteId).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("query row error: %w", err)
	}

	return &note, nil
}
