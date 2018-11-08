package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var readDB *sql.DB

func Initialize() error {
	readDB, err := sql.Open("sqlite3", "./content/data/bla.db")
	if err != nil {
		return err
	}

	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}

	if err = initializePosts(writeDB); err != nil {
		writeDB.Rollback()
		return err
	}

	if err = initializeUsers(writeDB); err != nil {
		writeDB.Rollback()
		return err
	}

	if err = initializeImages(writeDB); err != nil {
		writeDB.Rollback()
		return err
	}

	return writeDB.Commit()
}

func initializePosts(db *sql.Tx) error {
	statement := `
		CREATE TABLE IF NOT EXISTS posts (
			id BLOB PRIMARY KEY,
			tag TEXT,
  			title TEXT,
  			content_md TEXT,
			content_html TEXT,
			published TIMESTAMP,
			edited TIMESTAMP,
			is_favorite BOOLEAN,
  			author BLOB
		);
	`

	_, err := db.Exec(statement)

	return err
}

func initializeUsers(db *sql.Tx) error {
	statement := `
		CREATE TABLE IF NOT EXISTS users (
			id BLOB PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			email TEXT,
			location TEXT,
			catch_phrase TEXT,
			login TEXT,
			hashed_pw TEXT,
			salt TEXT	
		);
	`

	_, err := db.Exec(statement)

	return err
}

func initializeImages(db *sql.Tx) error {
	statement := `
		CREATE TABLE IF NOT EXISTS images (
			id BLOB PRIMARY KEY,
			uploaded TIMESTAMP,
			path TEXT	
		);
	`

	_, err := db.Exec(statement)

	return err
}
