package models

type User struct {
	ID           string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
