package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {

	dir, err := filepath.Abs("./src/api/db/")
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(dir, "database.db")
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}
	return db, nil
}
