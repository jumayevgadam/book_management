package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/models"
)

type BookRepository struct {
	DB *pgxpool.Pool
}

func NewBookRepository(DB *pgxpool.Pool) *BookRepository {
	return &BookRepository{DB: DB}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	var exists bool
	query := `SELECT EXISTS(
				 SELECT 1 FROM authors 
				 WHERE id = $1)`

	err := r.DB.QueryRow(ctx, query, book.Author_ID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("failed in author check: %v", err.Error())
	}

	if !exists {
		return nil, fmt.Errorf("author with id %d does not exist", book.Author_ID)
	}

	query2 := `INSERT INTO books(
					title, author_id, year, genre)
					VALUES ($1, $2, $3, $4) 
					RETURNING id`

	err = r.DB.QueryRow(ctx, query2, book.Title, book.Author_ID, book.Year, book.Genre).Scan(&book.ID)
	if err != nil {
		return nil, fmt.Errorf("failed in book creation: %v", err.Error())
	}

	return book, nil
}

func (r *BookRepository) GetBookByID(ctx context.Context, book_id int) (*models.Book, error) {
	return nil, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.Book, error) {
	return nil, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error) {
	return "", nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, book_id int) (string, error) {
	return "", nil
}
