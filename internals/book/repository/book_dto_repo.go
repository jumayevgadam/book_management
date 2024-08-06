package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/pkg/logger"
	"github.com/labstack/echo/v4"
)

type IBookRepository interface {
	CreateBook(ctx echo.Context, book *models.BookDAO) (*models.BookDTO, error)
	GetBookByID(ctx echo.Context, book_id int) (*models.BookDTO, error)
	GetAllBooks(ctx echo.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error)
	UpdateBook(ctx echo.Context, book_id int, updateInput *models.UpdateInputBook) (string, error)
	DeleteBook(ctx echo.Context, book_id int) (string, error)
}

type Repository struct {
	IBookRepository
}

func NewDTORepository(DB *pgxpool.Pool, logger logger.Logger) *Repository {
	return &Repository{
		IBookRepository: NewBookRepository(DB, logger),
	}
}
