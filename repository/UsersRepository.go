package repository

import (
	"errors"
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

	stmt, err := repo.db.Prepare("insert into users (user_id, username, email, password_hash) values ($1, $2, $3, $4)")
	if err != nil {
		return errors.New("statement prepare error" + err.Error())
	}

	r, err := stmt.Exec(userId, username, email, hashedPassword)
	if err != nil {
		//(duplicate key value violates unique constraint "users_username_key") for username conflicts
		//(duplicate key value violates unique constraint "users_email_key") for email conflicts
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			return errors.New("username already exists, cannot insert user" + err.Error())
		} else if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return errors.New("email already exists, cannot insert user" + err.Error())
		}
		return errors.New("statement execute error" + err.Error())
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return errors.New("rows affected error" + err.Error())
	}

	if rowsAffected > 1 {
		return errors.New("too many rows affected in table, failed insert user function")
	} else if rowsAffected == 0 {
		return errors.New("no rows affected in table, user not inserted")
	} else {
		fmt.Println(username, "inserted to database")
		return nil
	}
}

func (repo *Repo) SelectUserByUsername(username string) (*models.User, error) {
	var user models.User

	stmt, err := repo.db.Prepare("select user_id, email, password_hash from users where username = $1")
	if err != nil {
		return nil, errors.New("statement prepare error" + err.Error())
	}

	rows, err := stmt.Query(username)
	if err != nil {
		return nil, errors.New("statement query error" + err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, errors.New("rows scan error" + err.Error())
		}
	}

	return &user, nil
}

func (repo *Repo) SelectUserByEmail(email string) (*models.User, error) {
	var user models.User

	stmt, err := repo.db.Prepare("select user_id, username, password_hash from users where email = $1")
	if err != nil {
		return nil, errors.New("statement prepare error" + err.Error())
	}

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, errors.New("statement query error" + err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash)
		if err != nil {
			return nil, errors.New("rows scan error" + err.Error())
		}
	}

	return &user, nil
}

func (repo *Repo) SelectUserById(userId string) (*models.User, error) {
	var user models.User

	stmt, err := repo.db.Prepare("select username, email, password_hash from users where user_id = $1")
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
