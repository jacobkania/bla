package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Image struct {
	Id       uuid.UUID  `json:"id"`
	Uploaded *time.Time `json:"uploaded"`
	Path     string     `json:"path"`
}
