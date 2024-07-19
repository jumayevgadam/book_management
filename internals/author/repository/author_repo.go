package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/sirupsen/logrus"
)

type AuthorRepository struct {
	DB *pgxpool.Pool
}

func NewAuthorRepository(DB *pgxpool.Pool) *AuthorRepository {
	return &AuthorRepository{DB: DB}
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	query := `INSERT INTO authors (
					name, biography, birthdate) 
					VALUES ($1, $2, $3) 
					RETURNING id`

	err := r.DB.QueryRow(ctx, query, author.Name, author.Biography, author.Birthdate).Scan(&author.ID)
	if err != nil {
		logrus.Errorf("error in creating author(repo): %v", err)
		return nil, err
	}

	return author, nil
}

func (r *AuthorRepository) GetAuthorByID(ctx context.Context, author_id int) (*models.Author, error) {
	var OneAuthor models.Author

	query := `SELECT 
					id, name, biography, birthdate 
					FROM authors 
					WHERE id = $1`
	err := pgxscan.Get(ctx, r.DB, &OneAuthor, query, author_id)
	if err != nil {
		logrus.Errorf("error in fetching one author: %v", err.Error())
		return nil, err
	}

	return &OneAuthor, nil
}

func (r *AuthorRepository) GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.Author, error) {
	var Authors []*models.Author

	// Base query
	query := `SELECT * FROM authors`

	var args []interface{}
	argId := 1

	if pagination.Criteria != "" {
		query += fmt.Sprintf(` WHERE name ILIKE 
										$%d`, argId)
		args = append(args, "%%"+pagination.Criteria+"%%")
		argId++
	}

	query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`,
		argId, argId+1)
	args = append(args, pagination.Limit, pagination.Offset)

	err := pgxscan.Select(ctx, r.DB, &Authors, query, args...)
	if err != nil {
		logrus.Errorf("error in fetching all authors: %v", err.Error())
		return nil, err
	}

	return Authors, nil
}

func (r *AuthorRepository) UpdateAuthor(ctx context.Context, author_id int, updateInput *models.UpdateInputAuthor) (string, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateInput.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argId))
		args = append(args, *updateInput.Name)
		argId++
	}

	if updateInput.Biography != nil {
		setValues = append(setValues, fmt.Sprintf("biography = $%d", argId))
		args = append(args, *updateInput.Biography)
		argId++
	}

	if updateInput.Birthdate != nil {
		setValues = append(setValues, fmt.Sprintf("birthdate = $%d", argId))
		args = append(args, *updateInput.Birthdate)
		argId++
	}

	if len(setValues) == 0 {
		return "", fmt.Errorf("no fields for update")
	}

	query := fmt.Sprintf(`UPDATE authors SET 
									%s WHERE id = $%d 
									RETURNING 'Author information updated'`,
		strings.Join(setValues, ", "), argId)
	args = append(args, author_id)

	var response string
	_, err := r.DB.Exec(ctx, query, args...)
	if err != nil {
		logrus.Errorf("error in updating author: %v", err.Error())
		return response, err
	}

	response = fmt.Sprintf("Author with ID %d updated successfully", author_id)
	return response, nil
}

func (r *AuthorRepository) DeleteAuthor(ctx context.Context, author_id int) (string, error) {
	query := `DELETE FROM authors 
					WHERE id = $1 
					RETURNING 'Author deleted'`
	var response string

	err := r.DB.QueryRow(ctx, query, author_id).Scan(&response)
	if err != nil {
		logrus.Errorf("error in deleting author: %v", err.Error())
		return response, err
	}

	response = fmt.Sprintf("Author with ID %d deleted successfully", author_id)
	return response, nil
}
