package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/pkg/logger"
)

type IAuthorRepository interface {
	CreateAuthor(ctx context.Context, author *models.AuthorDAO) (*models.AuthorDTO, error)
	GetAuthorByID(ctx context.Context, author_id int) (*models.AuthorDTO, error)
	GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error)
	UpdateAuthor(ctx context.Context, author_id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx context.Context, author_id int) (string, error)
}

type Repository struct {
	IAuthorRepository
}

func NewDTORepository(DB *pgxpool.Pool, logger logger.Logger) *Repository {
	return &Repository{
		IAuthorRepository: NewAuthorRepository(DB, logger),
	}
}
