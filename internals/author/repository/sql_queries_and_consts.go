package repository

const (
	// For handling caused error directory(easily)
	createAuthorDir  = "err in repository.(author_repo.go).CreateAuthor"
	getAuthorByIDDir = "err in repository.(author_repo.go).GetAuthorByID"
	getAllAuthorDir  = "err in repository.(author_repo.go).GetAllAuthor"
	updateAuthorDir  = "err in repository.(author_repo.go).UpdateAuthor"
	deleteAuthorDir  = "err in repository.(author_repo.go).DeleteAuthor"
)

const (
	createAuthorQuery = `INSERT INTO authors (
									name, biography, birthdate) 
									VALUES ($1, $2, $3) 
									RETURNING id`

	getOneAuthorQuery = `SELECT 
									id, name, biography, birthdate 
									FROM authors 
									WHERE id = $1`

	deleteAuthorQuery = `DELETE FROM authors 
									WHERE id = $1 
									RETURNING 'Author deleted'`
)
