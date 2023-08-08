package repository

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/google/uuid"
)

/*
	Create -> Insert
	Read   -> Select
	Update -> Update
	Delete -> Delete
*/

// InsertUser user to database table users (username, email ,hashed_password)
func (repo *Repo) InsertUser(username string, email string, hashedPassword string) error {

	// Generate uuid for user_id
	userId := uuid.New().String()

	stmt, err := repo.db.Prepare("INSERT INTO users (user_id, username, email, password_hash) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("statement prepare error %w", err)
	}

	_, err = stmt.Exec(userId, username, email, hashedPassword)
	if err != nil {
		//(duplicate key value violates unique constraint "users_username_key") for username conflicts
		//(duplicate key value violates unique constraint "users_email_key") for email conflicts
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			return fmt.Errorf("username already exists, cannot insert user %w", err)
		} else if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return fmt.Errorf("email already exists, cannot insert user %w", err)
		}
		return fmt.Errorf("statement execute error %w", err)
	}

	fmt.Println("user inserted to database")
	return nil
}

func (repo *Repo) SelectUserByUsername(username string) (*models.User, error) {
	var user models.User
	user.Username = username

	stmt, err := repo.db.Prepare("SELECT user_id, email, password_hash FROM users WHERE username = $1")
	if err != nil {
		return nil, fmt.Errorf("statement prepare error %w", err)
	}

	rows, err := stmt.Query(username)
	if err != nil {
		return nil, fmt.Errorf("statement query error %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, fmt.Errorf("rows scan error %w", err)
		}
	}

	return &user, nil
}

// SelectUserByEmail selects user from database table users by email
func (repo *Repo) SelectUserByEmail(email string) (*models.User, error) {
	var user models.User
	user.Email = email

	stmt, err := repo.db.Prepare("SELECT user_id, username, password_hash FROM users WHERE email = $1")
	if err != nil {
		return nil, fmt.Errorf("statement prepare error %w", err)
	}

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, fmt.Errorf("statement query error %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash)
		if err != nil {
			return nil, fmt.Errorf("rows scan error %w", err)
		}
	}

	return &user, nil
}

// SelectUserById selects user from database table users by user_id
func (repo *Repo) SelectUserById(userId string) (*models.User, error) {
	var user models.User

	stmt, err := repo.db.Prepare("SELECT username, email, password_hash FROM users WHERE user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("statement prepare error: %w", err)
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("statement query error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, fmt.Errorf("rows scan error: %w", err)
		}
	}

	return &user, nil
}

// UpdateUserPassword updates user password in database table users
func (repo *Repo) UpdateUserPassword(userId string, hashedPassword string) error {
	stmt, err := repo.db.Prepare("UPDATE users SET password_hash = $1 WHERE user_id = $2")
	if err != nil {
		return fmt.Errorf("statement prepare error while updating user password:" + err.Error())
	}

	_, err = stmt.Exec(hashedPassword, userId)
	if err != nil {
		return fmt.Errorf("statement execute error while updating user password:" + err.Error())
	}

	fmt.Println(userId, "password updated")
	return nil
}

// UpdateUserEmail updates user email in database table users
func (repo *Repo) UpdateUserEmail(userId, email string) error {
	stmt, err := repo.db.Prepare("UPDATE users SET email = $1 WHERE user_id = $2")
	if err != nil {
		return fmt.Errorf("statement prepare error while updating user email:" + err.Error())
	}

	_, err = stmt.Exec(email, userId)
	if err != nil {
		return fmt.Errorf("statement execute error while updating user email:" + err.Error())
	}

	fmt.Println(userId, "email updated")
	return nil
}

// DeleteUser deletes user from database table users
func (repo *Repo) DeleteUser(Id string) error {
	stmt, err := repo.db.Prepare("DELETE FROM users WHERE user_id = $1")
	if err != nil {
		return fmt.Errorf("statement prepare error while deleting user:" + err.Error())
	}

	_, err = stmt.Exec(Id)
	if err != nil {
		return fmt.Errorf("statement execute error while deleting user:" + err.Error())
	}

	fmt.Println(Id, "deleted from database")
	return nil
}
