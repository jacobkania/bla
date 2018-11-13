package models

import "github.com/gofrs/uuid"

// Contains all information about a user (admin) of the blog. The Login and Hashed Password
// information is never returned in json responses.
type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Location    string    `json:"location"`
	CatchPhrase string    `json:"catchPhrase"`
	Login       string    `json:"-"`
	HashedPw    string    `json:"-"`
}
