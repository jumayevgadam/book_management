package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/models"
)

// AuthorRepository Interface
type IAuthorRepository interface {
	CreateAuthor(ctx context.Context, author *models.AuthorDAO) (*models.AuthorDTO, error)
	GetAuthorByID(ctx context.Context, author_id int) (*models.AuthorDTO, error)
	GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error)
	UpdateAuthor(ctx context.Context, author_id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx context.Context, author_id int) (string, error)
}

// Repository is
type Repository struct {
	IAuthorRepository
}

// New Data Transfer Repository is
func NewDTORepository(DB *pgxpool.Pool) *Repository {
	return &Repository{
		IAuthorRepository: NewAuthorRepository(DB),
	}
}
