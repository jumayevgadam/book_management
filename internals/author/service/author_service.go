package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/repository"
)

// AuthorService is
type AuthorService struct {
	repo repository.IAuthorRepository
}

// NewAuthorService is
func NewAuthorService(repo *repository.IAuthorRepository) *AuthorService {
	return &AuthorService{repo: *repo}
}

// CreateAuthor Service is
func (s *AuthorService) CreateAuthor(ctx context.Context, author *models.AuthorDAO) (*models.AuthorDTO, error) {
	return s.repo.CreateAuthor(ctx, author)
}

// GetAuthorByID Service is
func (s *AuthorService) GetAuthorByID(ctx context.Context, author_id int) (*models.AuthorDTO, error) {
	return s.repo.GetAuthorByID(ctx, author_id)
}

// GetAllAuthor Service is
func (s *AuthorService) GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error) {
	return s.repo.GetAllAuthor(ctx, pagination)
}

// UpdateAuthor Service is
func (s *AuthorService) UpdateAuthor(ctx context.Context, id int, updateInput *models.UpdateInputAuthor) (string, error) {
	return s.repo.UpdateAuthor(ctx, id, updateInput)
}

// DeleteAuthor Service is
func (s *AuthorService) DeleteAuthor(ctx context.Context, author_id int) (string, error) {
	return s.repo.DeleteAuthor(ctx, author_id)
}
