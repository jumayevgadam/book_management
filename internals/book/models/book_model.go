package models

/*
	In this models package, we need informations about
	Book, UpdateInput and Filter
	(Maybe, bad documentation and bad note :)
*/

type BookDTO struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author_ID int    `json:"author_id"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
}

type BookDAO struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Author_ID int    `db:"author_id"`
	Year      int    `db:"year"`
	Genre     string `db:"genre"`
}

type UpdateInputBook struct {
	Title *string `json:"title"`
	Year  *int    `json:"year"`
	Genre *string `json:"genre"`
}

type PaginationForBook struct {
	Limit  int    `json:"limit" form:"limit" binding:"required"`
	Offset int    `json:"offset" form:"offset" binding:"required"`
	Title  string `json:"title" form:"title"`
	Year   int    `json:"year" form:"year"`
	Genre  string `json:"Genre" form:"genre"`
}
