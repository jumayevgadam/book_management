package models

/*
	In this models package, we need informations about
	Book, UpdateInput and Filter
	(Maybe, bad documentation and bad note :)
*/

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author_ID int    `json:"author_id"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
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
