package database

import (
	"database/sql"
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/config"
	"github.com/MelihEmreGuler/go-user-notes-app/repository"
)

var (
	DB    *sql.DB
	dbErr error
)

// Init function is called from main.go
func Init() {
	//Connect to database
	Connect()

	// Create a new repository (singleton pattern) (only one instance of this struct)
	repository.NewRepo(DB)

	// Create tables if not exists to database
	CreateTables()
}

func Connect() {

	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	) // Sprintf returns string

	//open database
	open(psqlInfo)
}

func open(psqlInfo string) {
	DB, dbErr = sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		fmt.Println("database open err:", dbErr)
		return
	}

	err := DB.Ping()
	if err != nil {
		fmt.Println("database ping err:", err)
		return
	}

	fmt.Println("database open success")
}
