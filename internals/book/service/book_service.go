package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/repository"
)

type BookService struct {
	repo repository.IBookRepository
}

func NewBookService(repo *repository.IBookRepository) *BookService {
	return &BookService{repo: *repo}
}

func (s *BookService) CreateBook(ctx context.Context, book *models.BookDAO) (*models.BookDTO, error) {
	return s.repo.CreateBook(ctx, book)
}

func (s *BookService) GetBookByID(ctx context.Context, book_id int) (*models.BookDTO, error) {
	return s.repo.GetBookByID(ctx, book_id)
}

func (s *BookService) GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error) {
	// transction idea
	return s.repo.GetAllBooks(ctx, pagination)
}

func (s *BookService) UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error) {
	return s.repo.UpdateBook(ctx, book_id, updateInput)
}

func (s *BookService) DeleteBook(ctx context.Context, book_id int) (string, error) {
	return s.repo.DeleteBook(ctx, book_id)
}
