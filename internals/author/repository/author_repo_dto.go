package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/pkg/logger"
	"github.com/labstack/echo/v4"
)

type IAuthorRepository interface {
	CreateAuthor(ctx echo.Context, author *models.AuthorDAO) (*models.AuthorDTO, error)
	GetAuthorByID(ctx echo.Context, author_id int) (*models.AuthorDTO, error)
	GetAllAuthor(ctx echo.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error)
	UpdateAuthor(ctx echo.Context, author_id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx echo.Context, author_id int) (string, error)
}

type Repository struct {
	IAuthorRepository
}

func NewDTORepository(DB *pgxpool.Pool, logger logger.Logger) *Repository {
	return &Repository{
		IAuthorRepository: NewAuthorRepository(DB, logger),
	}
}
