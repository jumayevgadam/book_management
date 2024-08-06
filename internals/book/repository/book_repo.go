package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/pkg/logger"
	"github.com/labstack/echo/v4"
)

type BookRepository struct {
	DB     *pgxpool.Pool
	logger logger.Logger
}

func NewBookRepository(DB *pgxpool.Pool, logger logger.Logger) *BookRepository {
	return &BookRepository{
		DB:     DB,
		logger: logger,
	}
}

func (r *BookRepository) CreateBook(c echo.Context, book *models.BookDAO) (*models.BookDTO, error) {
	ctx := c.Request().Context()

	pgTx, err := r.DB.Begin(ctx)
	if err != nil {
		r.logger.Errorf("failed to begin transaction: %v", err.Error())
		return nil, err
	}

	var exists bool
	err = pgTx.QueryRow(ctx, existanceAuthorIDQuery, book.Author_ID).Scan(&exists)
	if err != nil {
		r.logger.Errorf("invalid author id")
		return nil, fmt.Errorf("failed in author check: %v", err.Error())
	}

	if !exists {
		r.logger.Errorf("does not exist author with that id")
		return nil, fmt.Errorf("author with id %d does not exist", book.Author_ID)
	}

	if book.Year > time.Now().Year() {
		r.logger.Errorf("invalid year %d", book.Year)
		return nil, fmt.Errorf("invalid year %d", book.Year)
	}

	err = pgTx.QueryRow(ctx, createBookQuery, book.Title, book.Author_ID, book.Year, book.Genre).Scan(&book.ID)
	if err != nil {
		r.logger.Errorf("error in book creation: %v", err.Error())
		return nil, err
	}

	err = pgTx.Commit(ctx)
	if err != nil {
		r.logger.Errorf("failed to commit transaction: %v", err.Error())
		return nil, err
	}

	return models.ConvertBookDAOToDTO(book), nil
}

func (r *BookRepository) GetBookByID(c echo.Context, book_id int) (*models.BookDTO, error) {
	ctx := c.Request().Context()

	var Book models.BookDAO
	err := pgxscan.Get(ctx, r.DB, &Book, gettingOneBookQuery, book_id)
	if err != nil {
		r.logger.Errorf("error in fetching one book: %v", err.Error())
		return nil, err
	}

	return models.ConvertBookDAOToDTO(&Book), nil
}

// Author can be search books about with title or published year
// Paginnation also need; generally filter need::
func (r *BookRepository) GetAllBooks(c echo.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error) {
	ctx := c.Request().Context()

	var Books []*models.BookDAO
	// Base query
	query := `SELECT * FROM books`

	var args []interface{}
	argId := 1
	conditions := []string{}

	if pagination.Title != "" {
		conditions = append(conditions, fmt.Sprintf("title ILIKE $%d", argId))
		args = append(args, fmt.Sprintf("%%%s%%", pagination.Title))
		argId++
	}

	if pagination.Year != 0 {
		conditions = append(conditions, fmt.Sprintf("year = $%d", argId))
		args = append(args, pagination.Year)
		argId++
	}

	if pagination.Genre != "" {
		conditions = append(conditions, fmt.Sprintf("genre ILIKE $%d", argId))
		args = append(args, fmt.Sprintf("%%%s%%", pagination.Genre))
		argId++
	}

	if len(conditions) > 0 {
		query += " WHERE " + fmt.Sprintf(" %s", conditions[0])
		for i := 1; i < len(conditions); i++ {
			query += " AND " + conditions[i]
		}
	}

	query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`,
		argId, argId+1)
	args = append(args, pagination.Limit, pagination.Offset)

	err := pgxscan.Select(ctx, r.DB, &Books, query, args...)
	if err != nil {
		r.logger.Errorf("error in fetching all books: %v", err.Error())
		return nil, err
	}

	var BookDTOs []*models.BookDTO
	for _, book := range Books {
		BookDTOs = append(BookDTOs, models.ConvertBookDAOToDTO(book))
	}

	return BookDTOs, nil
}

func (r *BookRepository) UpdateBook(c echo.Context, book_id int, updateInput *models.UpdateInputBook) (string, error) {
	ctx := c.Request().Context()

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *updateInput.Title)
		argId++
	}

	if updateInput.Year != nil {
		setValues = append(setValues, fmt.Sprintf("year = $%d", argId))
		args = append(args, *updateInput.Year)
		argId++
	}

	if updateInput.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre = $%d", argId))
		args = append(args, *updateInput.Genre)
		argId++
	}

	if len(setValues) == 0 {
		return "", fmt.Errorf("no fields for update")
	}

	query := fmt.Sprintf(`UPDATE books SET
								%s WHERE id = $%d
								RETURNING 'Book informations updated'
								`, strings.Join(setValues, ", "), argId)
	args = append(args, book_id)

	var response string
	_, err := r.DB.Exec(ctx, query, args...)
	if err != nil {
		r.logger.Errorf("error in updating book: %v", err.Error())
		return response, err
	}

	r.logger.Debugf("Updated query: ", query)

	response = fmt.Sprintf("Book with ID %d updated successfully", book_id)
	return response, nil
}

func (r *BookRepository) DeleteBook(c echo.Context, book_id int) (string, error) {
	ctx := c.Request().Context()

	var response string
	err := r.DB.QueryRow(ctx, deleteBookQuery, book_id).Scan(&response)
	if err != nil {
		r.logger.Errorf("error in deleting book: %v", err.Error())
		return response, err
	}

	response = fmt.Sprintf("Book with ID %d deleted successfully", book_id)
	return response, nil
}
