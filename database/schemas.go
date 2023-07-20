package database

import (
	"fmt"
)

func CreateTables() {
	createUsersTable()
	createNoteTable()
}

// CreateUsersTable func to create users table ---
func createUsersTable() {

	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(100) NOT NULL
);`)

	if err != nil {
		fmt.Printf("Error executing SQL query : %v\n", err)
		return
	}
}

// CreateNoteTable func to create notes table ---
func createNoteTable() {

	_, err := DB.Exec(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS notes (
    note_id UUID PRIMARY KEY,
	user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);`))

	if err != nil {
		fmt.Printf("Error executing SQL query: %v\n", err)
	}
}
