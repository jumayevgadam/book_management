package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/repository"
)

// AuthorService Interface
type IAuthorService interface {
	CreateAuthor(ctx context.Context, author *models.AuthorDAO) (*models.AuthorDTO, error)
	GetAuthorByID(ctx context.Context, author_id int) (*models.AuthorDTO, error)
	GetAllAuthor(ctx context.Context, pagination models.PaginationForAuthor) ([]*models.AuthorDTO, error)
	UpdateAuthor(ctx context.Context, author_id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx context.Context, author_id int) (string, error)
}

// Service struct is
type Service struct {
	IAuthorService
}

// New Data Transfer Service is
func NewDTOService(repo *repository.Repository) *Service {
	return &Service{
		IAuthorService: NewAuthorService(&repo.IAuthorRepository),
	}
}
