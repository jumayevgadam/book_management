package service

import (
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/repository"
	"github.com/labstack/echo/v4"
)

type IAuthorService interface {
	CreateAuthor(ctx echo.Context, author *models.AuthorDAO) (*models.AuthorDTO, error)
	GetAuthorByID(ctx echo.Context, author_id int) (*models.AuthorDTO, error)
	GetAllAuthor(ctx echo.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error)
	UpdateAuthor(ctx echo.Context, author_id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx echo.Context, author_id int) (string, error)
}

type Service struct {
	IAuthorService
}

func NewDTOService(repo *repository.Repository) *Service {
	return &Service{
		IAuthorService: NewAuthorService(&repo.IAuthorRepository),
	}
}
