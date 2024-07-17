package models

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author_ID int       `json:"author_id"`
	Year      time.Time `json:"year"`
	Genre     string    `json:"genre"`
}
