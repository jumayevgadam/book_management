package service

import (
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/repository"
	"github.com/labstack/echo/v4"
)

type BookService struct {
	repo repository.IBookRepository
}

func NewBookService(repo *repository.IBookRepository) *BookService {
	return &BookService{repo: *repo}
}

func (s *BookService) CreateBook(ctx echo.Context, book *models.BookDAO) (*models.BookDTO, error) {
	return s.repo.CreateBook(ctx, book)
}

func (s *BookService) GetBookByID(ctx echo.Context, book_id int) (*models.BookDTO, error) {
	return s.repo.GetBookByID(ctx, book_id)
}

func (s *BookService) GetAllBooks(ctx echo.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error) {
	// transction idea
	return s.repo.GetAllBooks(ctx, pagination)
}

func (s *BookService) UpdateBook(ctx echo.Context, book_id int, updateInput *models.UpdateInputBook) (string, error) {
	return s.repo.UpdateBook(ctx, book_id, updateInput)
}

func (s *BookService) DeleteBook(ctx echo.Context, book_id int) (string, error) {
	return s.repo.DeleteBook(ctx, book_id)
}
