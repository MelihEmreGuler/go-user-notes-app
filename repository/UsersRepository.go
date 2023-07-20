package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

var UserRepository = &UserRepo{}

func NewRepo(db *sql.DB) {
	UserRepository = &UserRepo{
		db: db,
	}
}

/*
	Create -> Insert
	Read   -> Select
	Update -> Update
	Delete -> Delete
*/

// Insert user to database table users (username, email ,hashed_password)
func (repo UserRepo) Insert(username string, email string, hashedPassword string) error {

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

func (repo UserRepo) SelectByUsername(username string) *models.User {
	var user models.User

	stmt, err := repo.db.Prepare("select user_id, email, password_hash from users where username = $1")
	if err != nil {
		fmt.Println("statement prepare error:", err)
		return nil
	}

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println("statement query error:", err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.PasswordHash)
		if err != nil {
			fmt.Println("rows scan error:", err)
			return nil
		}
	}

	return &user
}

func (repo UserRepo) SelectByEmail(email string) *models.User {
	var user models.User

	stmt, err := repo.db.Prepare("select user_id, username, password_hash from users where email = $1")
	if err != nil {
		fmt.Println("statement prepare error:", err)
		return nil
	}

	rows, err := stmt.Query(email)
	if err != nil {
		fmt.Println("statement query error:", err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash)
		if err != nil {
			fmt.Println("rows scan error:", err)
			return nil
		}
	}

	return &user
}
