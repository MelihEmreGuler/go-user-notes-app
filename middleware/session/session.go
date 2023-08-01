package session

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/middleware/cookie"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/MelihEmreGuler/go-user-notes-app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func InitSession() {
	go garbageCollector()
}

// CreateSession creates a new session and stores it in the database and sends the session ID to the user's computer
func CreateSession(c *fiber.Ctx, userID string) error {
	//generate uuid for session_id
	sessionID := uuid.New().String()

	// Insert session to database
	if err := repository.R.InsertSession(sessionID, userID, c.IP(), c.Get("User-Agent")); err != nil {
		return fmt.Errorf("error while inserting session: %w", err)
	}

	// Create a cookie to send the generated session ID to the user's computer
	cookie.SetSessionCookie(c, sessionID)

	return nil
}

// AuthenticateAndRefresh checks if the user is logged in and refreshes the session and returns the session
func AuthenticateAndRefresh(c *fiber.Ctx) (*models.Session, error) {
	// Read the session ID from the cookie stored on the user's computer
	sessionID := cookie.GetSessionCookie(c)

	// Check if session exists in database
	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting session: %w", err)
	}

	// Update session
	if err = repository.R.UpdateSession(sess.ID, c.IP(), c.Get("User-Agent")); err != nil {
		return nil, fmt.Errorf("error while updating session: %w", err)
	}

	return sess, nil
}

// DeleteSession deletes the session from database and invalidates the session ID cookie on the user's computer
func DeleteSession(c *fiber.Ctx) error {
	// Read the session ID from the cookie stored on the user's computer
	sessionID := cookie.GetSessionCookie(c)

	// Delete session from database
	if err := repository.R.DeleteSession(sessionID); err != nil {
		return fmt.Errorf("error while deleting session: %w", err)
	}

	// Invalidate the session ID cookie on the user's computer to delete it
	cookie.DeleteSessionCookie(c)

	return nil
}

func garbageCollector() {
	// We start a loop that will run continuously
	for {
		// We select all sessions from the database
		sessions, err := repository.R.SelectAllSessionsLastEvent()
		if err != nil {
			fmt.Printf("error while selecting all sessions: %v\n", err)
			time.Sleep(10 * time.Second) // In case of error, we will wait 10 seconds and try again
			continue
		}

		// We're taking the current time
		now := time.Now().Local()

		var sessionIdList []string

		// Loop through all sessions and delete the expired ones
		for _, sess := range sessions {
			// If the last operation of the session was done before 24 hours, we will select it in sessionIdList
			if now.Local().Sub(sess.LastEvent) > 24*time.Hour {
				sessionIdList = append(sessionIdList, sess.ID)
			}
		}

		if len(sessionIdList) > 0 {
			// Delete all sessions in sessionIdList
			if err = repository.R.DeleteSessions(sessionIdList); err != nil {
				fmt.Printf("error while deleting sessions: %v\n", err)
				time.Sleep(10 * time.Second) // In case of error, we will wait 10 seconds and try again
				continue
			}
		}

		// We wait 30 minutes and enter the loop again
		time.Sleep(30 * time.Minute)
	}
}
