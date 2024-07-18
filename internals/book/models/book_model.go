package models

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
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
