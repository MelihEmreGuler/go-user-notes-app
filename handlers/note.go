package handlers

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/middleware/session"
	"github.com/MelihEmreGuler/go-user-notes-app/repository"
	"github.com/gofiber/fiber/v2"
)

type createNoteRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	SessionID string `json:"session_id"`
}
type updateNoteRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	SessionID string `json:"session_id"`
}
type deleteNoteRequest struct {
	SessionID string `json:"session_id"`
}

// checkSession checks if the user is logged in and sets the session information to the context
func checkSession(c *fiber.Ctx) error {
	// Checks if the user is logged in
	sess, err := session.AuthenticateAndRefresh(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	// Set the session information to the context
	c.Locals("session", sess)

	return nil
}

/*
curl -X POST http://localhost:8080/note \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Note",
    "content": "This is the content of my first note."
  }'
*/

// CreateNote creates a new note
// @Summary Create a new note
// @Description Create a new note with the given title and content
// @Tags Notes
// @Accept json
// @Produce json
// @Param request body createNoteRequest true "Note Information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /note [post]
func CreateNote(c *fiber.Ctx) error {

	// Checks if the user is logged in
	// if err := checkSession(c); err != nil {
	//	return err
	// }

	// Get the session information from the context
	//sess, ok := c.Locals("session").(*models.Session)
	// if !ok {
	//	return fmt.Errorf("session not found")
	// }

	request := &createNoteRequest{}
	if err := c.BodyParser(request); err != nil {
		return fmt.Errorf("error while parsing request body: %w", err)
	}

	sessionID := request.SessionID

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	// Insert note to database
	err = repository.R.InsertNote(sess.UserID, request.Title, request.Content)
	if err != nil {
		return fmt.Errorf("error while inserting note: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Note created",
	})
}

// GetNotes returns all notes of the user
// @Summary Get all notes
// @Description Get all notes of the user
// @Tags Notes
// @Accept json
// @Produce json
// @Param session_id header string true "Session ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /notes [get]
func GetNotes(c *fiber.Ctx) error {

	sessionID := c.Get("session_id")

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	// Get notes from database
	notes, err := repository.R.GetNotes(sess.UserID)
	if err != nil {
		return fmt.Errorf("error while getting notes: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"notes":   notes,
	})
}

// GetNote returns a note of the user
// @Summary Get a note
// @Description Get a note of the user with the given id
// @Tags Notes
// @Accept json
// @Produce json
// @Param id query string true "Note ID"
// @Param session_id header string true "Session ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /note/:id [get]
func GetNote(c *fiber.Ctx) error {

	sessionID := c.Get("session_id")

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	note, err := repository.R.GetNoteByID(sess.UserID, c.Query("id"))
	if err != nil {
		return fmt.Errorf("error while getting note: %w", err)
	}

	if note == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "note not found",
		})
	}

	// Set the user id to the note
	note.UserId = sess.UserID

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"note":    note,
	})
}

// UpdateNote updates a note of the user
// @Summary Update a note
// @Description Update a note of the user with the given id
// @Tags Notes
// @Accept json
// @Produce json
// @Param id query string true "Note ID"
// @Param request body updateNoteRequest true "Note Information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /note/:id [put]
func UpdateNote(c *fiber.Ctx) error {

	request := &updateNoteRequest{}
	if err := c.BodyParser(request); err != nil {
		return fmt.Errorf("error while parsing request body: %w", err)
	}

	sessionID := request.SessionID

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	err = repository.R.UpdateNoteByID(sess.UserID, c.Query("id"), request.Title, request.Content)
	if err != nil {
		return fmt.Errorf("error while updating note: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Note updated",
	})
}

// DeleteNote deletes a note of the user
// @Summary Delete a note
// @Description Delete a note of the user with the given id
// @Tags Notes
// @Accept json
// @Produce json
// @Param id query string true "Note ID"
// @Param request body deleteNoteRequest true "Note Information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /note/:id [delete]
func DeleteNote(c *fiber.Ctx) error {

	request := &deleteNoteRequest{}
	if err := c.BodyParser(request); err != nil {
		return fmt.Errorf("error while parsing request body: %w", err)
	}

	sessionID := request.SessionID

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	err = repository.R.DeleteNoteByID(sess.UserID, c.Query("id"))
	if err != nil {
		return fmt.Errorf("error while deleting note: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "note deleted",
	})
}
