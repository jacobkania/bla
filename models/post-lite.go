package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// Contains only the information about a blog post that is necessary
// for displaying in a list of blog posts. ie. to be able to identify the post
type PostLite struct {
	Id         uuid.UUID  `json:"id"`
	Tag        string     `json:"tag"`
	Title      string     `json:"title"`
	Published  *time.Time `json:"published"`
	IsFavorite bool       `json:"isFavorite"`
}
