package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/repository"
)

type IBookService interface {
	CreateBook(ctx context.Context, book *models.BookDAO) (*models.BookDTO, error)
	GetBookByID(ctx context.Context, book_id int) (*models.BookDTO, error)
	GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error)
	UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error)
	DeleteBook(ctx context.Context, book_id int) (string, error)
}

type Service struct {
	IBookService
}

func NewDTOService(repo *repository.Repository) *Service {
	return &Service{
		IBookService: NewBookService(&repo.IBookRepository),
	}
}
