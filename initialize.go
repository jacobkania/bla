package main

import (
	"bla/authentication"
	"bla/storage"
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

// Runs the commands necessary for first-time setup of the application. This creates the necessary
// tables in the database if they don't exist, and prompts for admin user info if no user exists.
// It is completely safe to run this method when data and users already exist, as it is non-destructive.
func Initialize(db *sql.DB) error {
	if err := createDB(db); err != nil {
		return err
	}
	if err := promptUserInfo(db); err != nil {
		return err
	}
	return nil
}

// Calls to the Initialize method in the storage package, which will initialize the database tables.
func createDB(db *sql.DB) error {
	return storage.Initialize(db)
}

// If no user exists in the database, prompts for the necessary user information to populate
// an initial admin user. If any user already exists, does nothing.
func promptUserInfo(db *sql.DB) error {
	if count, _ := storage.GetUserCount(db); count != 0 {
		return nil
	}

	// get user info like username/password
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

	hashedPw, _ := authentication.EncryptPassword(password[:len(password)-1])

	newUser, err := storage.CreateUser(db, firstName[:len(firstName)-1], lastName[:len(lastName)-1], email[:len(email)-1], location[:len(location)-1], catchPhrase[:len(catchPhrase)-1], login[:len(login)-1], hashedPw)
	if err != nil {
		return err
	}

	fmt.Printf("User %s created successfully.", newUser.FirstName)

	return nil
}
