package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=gostudy dbname=gostudy password=gostudy sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
