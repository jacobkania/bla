package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// Contains information for an Image which will be embedded in a blog post.
// Path is the directory path that can be followed to reach the image
type Image struct {
	Id       uuid.UUID  `json:"id"`
	Uploaded *time.Time `json:"uploaded"`
	Path     string     `json:"path"`
}
