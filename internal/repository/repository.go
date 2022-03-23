package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(databaseURL string) *Repository {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}
	return &Repository{
		db: db,
	}
}
