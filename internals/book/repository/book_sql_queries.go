package repository

const (
	existanceAuthorIDQuery = `SELECT EXISTS(
										SELECT 1 FROM authors 
										WHERE id = $1)`

	createBookQuery = `INSERT INTO books(
								title, author_id, year, genre)
								VALUES ($1, $2, $3, $4) 
								RETURNING id`

	gettingOneBookQuery = `SELECT
									id, title, author_id, year, genre
									FROM books
									WHERE id = $1`

	deleteBookQuery = `DELETE FROM books 
								WHERE id = $1 
								RETURNING 'Book deleted'`
)
