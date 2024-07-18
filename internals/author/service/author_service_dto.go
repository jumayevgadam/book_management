package service

import (
	"context"

	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/repository"
)

type AuthorDTO interface {
	CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error)
	GetAuthorByID(ctx context.Context, author_id int) (*models.Author, error)
	GetAllAuthor(ctx context.Context) ([]*models.Author, error)
	UpdateAuthor(ctx context.Context, id int, update *models.UpdateInputAuthor) (string, error)
	DeleteAuthor(ctx context.Context, author_id int) (string, error)
}

type Service struct {
	AuthorDTO
}

func NewDTOService(repo *repository.Repository) *Service {
	return &Service{
		AuthorDTO: NewAuthorService(&repo.AuthorDTO),
	}
}
