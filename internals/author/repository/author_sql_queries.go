package repository

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
