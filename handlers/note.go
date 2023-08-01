package handlers

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/middleware/session"
	"github.com/MelihEmreGuler/go-user-notes-app/repository"
	"github.com/gofiber/fiber/v2"
)

/*
curl -X POST http://localhost:8080/note \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Note",
    "content": "This is the content of my first note."
  }'
*/

// CreateNote creates a new note
func CreateNote(c *fiber.Ctx) error {
	// Checks if the user is logged in
	sess, err := session.AuthenticateAndRefresh(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	type createNoteRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	request := &createNoteRequest{}
	if err = c.BodyParser(request); err != nil {
		return fmt.Errorf("error while parsing request body: %w", err)
	}

	// Insert note to database
	err = repository.R.InsertNote(sess.UserID, request.Title, request.Content)
	if err != nil {
		return fmt.Errorf("error while inserting note: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "note created",
	})
}

// GetNotes returns all notes of the user
