package initialize

import (
	"bla/authentication"
	"bla/storage"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

func Initialize(db *sql.DB) error {
	// do all of the setup for first time running the application
	if err := createDB(db); err != nil {
		return err
	}
	if err := promptUserInfo(db); err != nil {
		return err
	}
	if err := createConfig(); err != nil {
		return err
	}
	return nil
}

func createDB(db *sql.DB) error {
	return storage.Initialize(db)
}

func promptUserInfo(db *sql.DB) error {
	if count, _ := storage.GetUserCount(db); count != 0 {
		return nil
	}

	// get user info like owner username/password
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("No user currently exists! Please create one:\n")
	fmt.Printf("First name: ")
	firstName, _ := reader.ReadString('\n')
	fmt.Printf("Last Name: ")
	lastName, _ := reader.ReadString('\n')
	fmt.Printf("Email: ")
	email, _ := reader.ReadString('\n')
	fmt.Printf("Location: ")
	location, _ := reader.ReadString('\n')
	fmt.Printf("Catch Phrase: ")
	catchPhrase, _ := reader.ReadString('\n')
	fmt.Printf("Login: ")
	login, _ := reader.ReadString('\n')
	fmt.Printf("Password: ")
	password, _ := reader.ReadString('\n')

	hashedPw, _ := authentication.EncryptPassword(password)

	newUser, err := storage.CreateUser(db, firstName, lastName, email, location, catchPhrase, login, hashedPw)
	if err != nil {
		return err
	}

	fmt.Printf("User %s created successfully.", newUser.FirstName)

	return nil
}

func createConfig() error {
	// creates the config file for basic config info
	// in content/data/config.json

	return nil
}
