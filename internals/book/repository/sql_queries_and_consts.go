package repository

const (
	// For taking easily address and they are
	createBookDir  = "err in repository.book_repo.CreateBook"
	getAllBooksDir = "err in repository.book_repo.GetAllBooks"
	getBookByIDDir = "err in repository.book_repo.GetBookByID"
	updateBookDir  = "err in repository.book_repo.UpdateBook"
	deleteBookDir  = "err in repository.book_repo.DeleteBook"
)

const (
	// Queries are
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
