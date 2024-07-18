package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/repository"
)

type BookDTO interface {
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	GetBookByID(ctx context.Context, book_id int) (*models.Book, error)
	GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.Book, error)
	UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error)
	DeleteBook(ctx context.Context, book_id int) (string, error)
}

type Service struct {
	BookDTO
}

func NewDTOService(repo *repository.Repository) *Service {
	return &Service{
		BookDTO: NewBookService(&repo.BookDTO),
	}
}
