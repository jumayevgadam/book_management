package models

import "time"

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Biography string    `json:"biography"`
	Birthdate time.Time `json:"birthdate"`
}

type UpdateInputAuthor struct {
	Name      *string    `json:"name"`
	Biography *string    `json:"biography"`
	Birthdate *time.Time `json:"birthdate"`
}

type PaginationForAuthor struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
