package repository

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"strings"
	"time"
)

// InsertSession inserts a new Session to database
func (repo *Repo) InsertSession(sessionID, userID, IPAddr, userAgent string) error {

	//generate time for login_time and last_event
	now := time.Now()

	stmt, err := repo.db.Prepare("INSERT INTO sessions (session_id, user_id, login_time, last_event, ip_address, user_agent) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	_, err = stmt.Exec(sessionID, userID, now, now, IPAddr, userAgent)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	fmt.Println("Session inserted to database")
	return nil
}

// UpdateSession updates last_event, ip_address and user_agent of a Session
func (repo *Repo) UpdateSession(sessionID, IPAddr, userAgent string) error {
	//generate time for last_event
	now := time.Now()

	stmt, err := repo.db.Prepare("UPDATE sessions SET last_event = $1, ip_address = $2, user_agent = $3 WHERE session_id = $4")
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	_, err = stmt.Exec(now, IPAddr, userAgent, sessionID)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	fmt.Println("Session updated")
	return nil
}

// SelectSession selects a Session from database
func (repo *Repo) SelectSession(sessionID string) (*models.Session, error) {
	var sess models.Session

	stmt, err := repo.db.Prepare("SELECT session_id, user_id, login_time, last_event, ip_address, user_agent FROM sessions WHERE session_id = $1")
	if err != nil {
		return nil, fmt.Errorf("statement prepare error: %w", err)
	}

	err = stmt.QueryRow(sessionID).Scan(&sess.ID, &sess.UserID, &sess.LoginTime, &sess.LastEvent, &sess.IPAddr, &sess.UserAgent)
	if err != nil {
		return nil, fmt.Errorf("query row error: %w", err)
	}

	return &sess, nil
}

func (repo *Repo) SelectAllSessions() ([]models.Session, error) {
	var sessList []models.Session

	rows, err := repo.db.Query("SELECT session_id, user_id, login_time, last_event, ip_address, user_agent FROM sessions")
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	for rows.Next() {
		var sess models.Session
		err = rows.Scan(&sess.ID, &sess.UserID, &sess.LoginTime, &sess.LastEvent, &sess.IPAddr, &sess.UserAgent)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		sessList = append(sessList, sess)
	}

	return sessList, nil
}

// SelectAllSessionsLastEvent selects all Sessions from database with just last_event and session_id
func (repo *Repo) SelectAllSessionsLastEvent() ([]models.Session, error) {
	var sessList []models.Session
	rows, err := repo.db.Query("SELECT session_id, last_event FROM sessions")
	if err != nil {
		fmt.Println("query error: ", err)
	}

	for rows.Next() {
		var sess models.Session
		err = rows.Scan(&sess.ID, &sess.LastEvent)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		sessList = append(sessList, sess)
	}

	return sessList, nil
}

// DeleteSession deletes a Session from database
func (repo *Repo) DeleteSession(sessionID string) error {
	stmt, err := repo.db.Prepare("DELETE FROM sessions WHERE session_id = $1")
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	_, err = stmt.Exec(sessionID)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	fmt.Println("Session deleted")
	return nil
}

// DeleteSessions deletes multiple Sessions from database
func (repo *Repo) DeleteSessions(sessionIdList []string) error {

	// Create placeholders for session IDs
	placeholders := make([]string, len(sessionIdList))
	for i := range sessionIdList {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	// Create the SQL query with placeholders
	query := "DELETE FROM sessions WHERE session_id IN (" + strings.Join(placeholders, ",") + ")"

	// Prepare the statement with the query
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("statement prepare error: %w", err)
	}

	// Convert sessionIdList to interface{} slice to pass to Exec function
	args := make([]interface{}, len(sessionIdList))
	for i, id := range sessionIdList {
		args[i] = id
	}

	// Execute the statement with the sessionIdList as arguments
	_, err = stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}

	fmt.Println("Sessions deleted")
	return nil
}
