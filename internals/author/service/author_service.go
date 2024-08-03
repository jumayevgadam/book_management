package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/repository"
)

type AuthorService struct {
	repo repository.IAuthorRepository
}

func NewAuthorService(repo *repository.IAuthorRepository) *AuthorService {
	return &AuthorService{repo: *repo}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author *models.AuthorDAO) (*models.AuthorDTO, error) {
	return s.repo.CreateAuthor(ctx, author)
}

func (s *AuthorService) GetAuthorByID(ctx context.Context, author_id int) (*models.AuthorDTO, error) {
	return s.repo.GetAuthorByID(ctx, author_id)
}

func (s *AuthorService) GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error) {
	return s.repo.GetAllAuthor(ctx, pagination)
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, id int, updateInput *models.UpdateInputAuthor) (string, error) {
	return s.repo.UpdateAuthor(ctx, id, updateInput)
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, author_id int) (string, error) {
	return s.repo.DeleteAuthor(ctx, author_id)
}
