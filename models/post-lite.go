package models

import "time"

type PostLite struct {
	Tag       string     `json:"tag"`
	Title     string     `json:"title"`
	Published *time.Time `json:"published"`
}
