package server

import "time"

type JsonPost struct {
	Id         int64
	Title      string
	Markdown   string
	Html       string
	IsFavorite bool
	Published  *time.Time
	Edited     *time.Time
}
