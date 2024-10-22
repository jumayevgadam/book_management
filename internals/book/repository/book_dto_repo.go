package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/models"
)

type IBookRepository interface {
	CreateBook(ctx context.Context, book *models.BookDAO) (*models.BookDTO, error)
	GetBookByID(ctx context.Context, book_id int) (*models.BookDTO, error)
	GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error)
	UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error)
	DeleteBook(ctx context.Context, book_id int) (string, error)
}

type Repository struct {
	IBookRepository
}

func NewDTORepository(DB *pgxpool.Pool) *Repository {
	return &Repository{
		IBookRepository: NewBookRepository(DB),
	}
}
