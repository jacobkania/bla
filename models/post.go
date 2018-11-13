package models

import (
	"github.com/gofrs/uuid"
	"time"
)

// Contains all information about a blog post
type Post struct {
	Id          uuid.UUID  `json:"id"`
	Tag         string     `json:"tag"`
	Title       string     `json:"title"`
	ContentMD   string     `json:"contentMd"`
	ContentHTML string     `json:"contentHtml"`
	Published   *time.Time `json:"published"`
	Edited      *time.Time `json:"edited"`
	IsFavorite  bool       `json:"isFavorite"`
	Author      uuid.UUID  `json:"author"`
}
