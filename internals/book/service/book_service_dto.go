package service

import (
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/repository"
	"github.com/labstack/echo/v4"
)

type IBookService interface {
	CreateBook(ctx echo.Context, book *models.BookDAO) (*models.BookDTO, error)
	GetBookByID(ctx echo.Context, book_id int) (*models.BookDTO, error)
	GetAllBooks(ctx echo.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error)
	UpdateBook(ctx echo.Context, book_id int, updateInput *models.UpdateInputBook) (string, error)
	DeleteBook(ctx echo.Context, book_id int) (string, error)
}

type Service struct {
	IBookService
}

func NewDTOService(repo *repository.Repository) *Service {
	return &Service{
		IBookService: NewBookService(&repo.IBookRepository),
	}
}
