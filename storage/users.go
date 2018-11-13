package storage

import (
	"bla/models"
	"database/sql"
	"github.com/gofrs/uuid"
)

const sqlGetUserCount string = `SELECT COUNT(*) FROM users`
const sqlGetAllUsers string = `SELECT id, first_name, last_name, email, location, catch_phrase FROM users`
const sqlGetUserById string = `SELECT id, first_name, last_name, email, location, catch_phrase FROM users WHERE id = ?`
const sqlGetUserByPersonalLogin string = `SELECT id, first_name, last_name, email, location, catch_phrase, login, hashed_pw FROM users WHERE login = ?`
const sqlCreateUser string = `INSERT INTO users (id, first_name, last_name, email, location, catch_phrase, login, hashed_pw) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// GetUserCount returns the number of admin users which exist in the supplied database.
func GetUserCount(db *sql.DB) (int, error) {
	row := db.QueryRow(sqlGetUserCount)

	var count int

	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetAllUsers returns a slice containing the User models of all users in the supplied database.
// The Login and Hashed Password fields are ignored.
func GetAllUsers(db *sql.DB) (*[]models.User, error) {
	rows, err := db.Query(sqlGetAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Location, &user.CatchPhrase)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

// GetUserById returns a User model containing information about the given user in the supplied database.
// The Login and Hashed Password fields are ignored.
func GetUserById(db *sql.DB, id uuid.UUID) (*models.User, error) {
	row := db.QueryRow(sqlGetUserById, id)

	user := models.User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Location, &user.CatchPhrase)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByPersonalLogin returns a User model containing all information for the user in
// the supplied database. This includes the user's Login and Hashed Password.
// This method is meant to be used to retrieve information for a user attempting to authenticate with the system.
func GetUserByPersonalLogin(db *sql.DB, login string) (*models.User, error) {
	row := db.QueryRow(sqlGetUserByPersonalLogin, login)

	user := models.User{}

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Location, &user.CatchPhrase, &user.Login, &user.HashedPw)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a User object in the supplied database with the given information.
func CreateUser(db *sql.DB, firstName, lastName, email, location, catchPhrase, login, hashedPw string) (*models.User, error) {
	writeDB, err := db.Begin()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_id, err := uuid.NewV4()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_, err = writeDB.Exec(sqlCreateUser, _id, firstName, lastName, email, location, catchPhrase, login, hashedPw)
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	if err = writeDB.Commit(); err != nil {
		return nil, err
	}

	return GetUserById(db, _id)
}
