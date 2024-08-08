package repository

import (
	"fmt"
	"strings"

	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/models"
)

// AuthorRepository is
type AuthorRepository struct {
	DB *pgxpool.Pool
}

// NewAuthorRepository is
func NewAuthorRepository(DB *pgxpool.Pool) *AuthorRepository {
	return &AuthorRepository{
		DB: DB,
	}
}

// CreateAuthor Repo is
func (r *AuthorRepository) CreateAuthor(ctx context.Context, author *models.AuthorDAO) (*models.AuthorDTO, error) {
	pgTx, err := r.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%v.BeginTransaction]", createAuthorDir)
	}

	defer func() {
		if err != nil {
			pgTx.Rollback(ctx)
		} else {
			pgTx.Commit(ctx)
		}
	}()

	err = pgTx.QueryRow(
		ctx,
		createAuthorQuery,
		author.Name,
		author.Biography,
		author.Birthdate,
	).Scan(&author.ID)

	if err != nil {
		return nil, fmt.Errorf("[%v.QueryRow][%v]", createAuthorDir, err)
	}

	return models.ConvertAuthorDAOToDTO(author), nil
}

// Getting Author by ID is
func (r *AuthorRepository) GetAuthorByID(ctx context.Context, author_id int) (*models.AuthorDTO, error) {
	var OneAuthor models.AuthorDAO
	err := pgxscan.Get(
		ctx,
		r.DB,
		&OneAuthor,
		getOneAuthorQuery,
		author_id,
	)
	if err != nil {
		return nil, fmt.Errorf("[%v.Get][%v]", getAuthorByIDDir, err)
	}

	return models.ConvertAuthorDAOToDTO(&OneAuthor), nil
}

// GetAllAuthor is
func (r *AuthorRepository) GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error) {
	var Authors []*models.AuthorDAO
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

	err := pgxscan.Select(
		ctx,
		r.DB,
		&Authors,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("[%v.Select][%v]", getAllAuthorDir, err)
	}

	var AuthorDTOs []*models.AuthorDTO
	for _, author := range Authors {
		AuthorDTOs = append(AuthorDTOs, models.ConvertAuthorDAOToDTO(author))
	}

	return AuthorDTOs, nil
}

// UpdateAuthor is, we perform multiple update in this function
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
		return "", fmt.Errorf("[%v.checkingLenSetValues]", updateAuthorDir)
	}

	query := fmt.Sprintf(`UPDATE authors SET 
									%s WHERE id = $%d 
									RETURNING 'Author information updated'`,
		strings.Join(setValues, ", "), argId)
	args = append(args, author_id)

	var response string
	_, err := r.DB.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return "", fmt.Errorf("[%v.Exec][%v]", updateAuthorDir, err)
	}

	response = fmt.Sprintf("Author with ID %d updated successfully", author_id)
	return response, nil
}

// DeleteAuthor is
func (r *AuthorRepository) DeleteAuthor(ctx context.Context, author_id int) (string, error) {
	var response string

	err := r.DB.QueryRow(
		ctx,
		deleteAuthorQuery,
		author_id).Scan(&response)
	if err != nil {
		return response, fmt.Errorf("[%v.QueryRow][%v]", deleteAuthorDir, err)
	}

	response = fmt.Sprintf("Author with ID %d deleted successfully", author_id)
	return response, nil
}
