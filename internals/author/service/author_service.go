package service

import (
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/repository"
	"github.com/labstack/echo/v4"
)

type AuthorService struct {
	repo repository.IAuthorRepository
}

func NewAuthorService(repo *repository.IAuthorRepository) *AuthorService {
	return &AuthorService{repo: *repo}
}

func (s *AuthorService) CreateAuthor(ctx echo.Context, author *models.AuthorDAO) (*models.AuthorDTO, error) {
	return s.repo.CreateAuthor(ctx, author)
}

func (s *AuthorService) GetAuthorByID(ctx echo.Context, author_id int) (*models.AuthorDTO, error) {
	return s.repo.GetAuthorByID(ctx, author_id)
}

func (s *AuthorService) GetAllAuthor(ctx echo.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error) {
	return s.repo.GetAllAuthor(ctx, pagination)
}

func (s *AuthorService) UpdateAuthor(ctx echo.Context, id int, updateInput *models.UpdateInputAuthor) (string, error) {
	return s.repo.UpdateAuthor(ctx, id, updateInput)
}

func (s *AuthorService) DeleteAuthor(ctx echo.Context, author_id int) (string, error) {
	return s.repo.DeleteAuthor(ctx, author_id)
}
