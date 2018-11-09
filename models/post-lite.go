package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type PostLite struct {
	Id         uuid.UUID  `json:"id"`
	Tag        string     `json:"tag"`
	Title      string     `json:"title"`
	Published  *time.Time `json:"published"`
	IsFavorite bool       `json:"isFavorite"`
}
