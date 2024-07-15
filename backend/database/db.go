package database

import (
	"database/sql"

	_ "modernc.org/sqlite" // This is the pure Go SQLite driver
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        original_url TEXT NOT NULL,
        short_url TEXT NOT NULL UNIQUE
    )`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
