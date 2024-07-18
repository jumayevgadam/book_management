package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/models"
)

type AuthorDTO interface {
	CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error)
	GetAuthorByID(ctx context.Context, author_id int) (*models.Author, error)
	GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.Author, error)
	UpdateAuthor(ctx context.Context, author_id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx context.Context, author_id int) (string, error)
}

type Repository struct {
	AuthorDTO
}

func NewDTORepository(DB *pgxpool.Pool) *Repository {
	return &Repository{
		AuthorDTO: NewAuthorRepository(DB),
	}
}
