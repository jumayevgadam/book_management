package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/pkg/logger"
)

type BookDTO interface {
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	GetBookByID(ctx context.Context, book_id int) (*models.Book, error)
	GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.Book, error)
	UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error)
	DeleteBook(ctx context.Context, book_id int) (string, error)
}

type Repository struct {
	BookDTO
}

func NewDTORepository(DB *pgxpool.Pool, logger logger.Logger) *Repository {
	return &Repository{
		BookDTO: NewBookRepository(DB, logger),
	}
}
