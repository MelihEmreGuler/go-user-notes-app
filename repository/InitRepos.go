package repository

import "database/sql"

type Repo struct {
	db *sql.DB
}

var R = &Repo{}

func NewRepo(db *sql.DB) {
	R = &Repo{
		db: db,
	}
}
