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
	readDB.SetMaxIdleConns(256)
	err = readDB.Ping()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if err = initializePosts(readDB); err != nil {
		return err
	}

	if err = initializeUsers(readDB); err != nil {
		return err
	}

	if err = initializeImages(readDB); err != nil {
		return err
	}

	return nil
}

func initializePosts(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS posts (
			id BLOB NOT NULL PRIMARY KEY,
			tag TEXT,
  			title TEXT,
  			content_md TEXT,
			content_html TEXT,
			published TIMESTAMP,
			edited TIMESTAMP,
			is_favorite INTEGER,
  			author BLOB
		);
	`

	_, err := db.Exec(statement)

	return err
}

func initializeUsers(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS users (
			id BLOB NOT NULL PRIMARY KEY,
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

func initializeImages(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS images (
			id BLOB NOT NULL PRIMARY KEY,
			uploaded TIMESTAMP,
			path TEXT	
		);
	`

	_, err := db.Exec(statement)

	return err
}
