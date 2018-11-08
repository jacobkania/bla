package models

import "github.com/gofrs/uuid"

type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Location    string    `json:"location"`
	CatchPhrase string    `json:"catchPhrase"`
	Login       string    `json:"-"`
	HashedPw    string    `json:"-"`
	Salt        string    `json:"-"`
}
