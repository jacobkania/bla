package initialize

import (
	"bla/storage"
	"database/sql"
)

func Initialize(db *sql.DB) error {
	// do all of the setup for first time running the application
	if err := createDB(db); err != nil {
		return err
	}
	if err := createConfig(); err != nil {
		return err
	}
	if err := promptUserInfo(); err != nil {
		return err
	}
	return nil
}

func createDB(db *sql.DB) error {
	return storage.Initialize(db)
}

func createConfig() error {
	// creates the config file for basic config info
	// in content/data/config.json

	return nil
}

func promptUserInfo() error {
	// get user info like owner username/password

	return nil
}
