package models

import "time"

type AuthorDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Biography string    `json:"biography"`
	Birthdate time.Time `json:"birthdate"`
}

type AuthorDAO struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Biography string    `db:"biography"`
	Birthdate time.Time `db:"birthdate"`
}

type UpdateInputAuthor struct {
	Name      *string    `json:"name"`
	Biography *string    `json:"biography"`
	Birthdate *time.Time `json:"birthdate"`
}

type PaginationForAuthor struct {
	Limit    int    `json:"limit" form:"limit" binding:"required"`
	Offset   int    `json:"offset" form:"offset" binding:"required"`
	Criteria string `json:"criteria" form:"criteria"`
}
