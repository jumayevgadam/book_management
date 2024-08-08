package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/models"
	"golang.org/x/net/context"
)

type BookRepository struct {
	DB *pgxpool.Pool
}

func NewBookRepository(DB *pgxpool.Pool) *BookRepository {
	return &BookRepository{
		DB: DB,
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *models.BookDAO) (*models.BookDTO, error) {
	pgTx, err := r.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var exists bool
	err = pgTx.QueryRow(ctx, existanceAuthorIDQuery, book.Author_ID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("[%v.QueryRow]=[%v]", createBookDir, err)
	}

	if !exists {
		return nil, fmt.Errorf("[%v.ifNotExist]=[%d]", createBookDir, book.Author_ID)
	}

	if book.Year > time.Now().Year() {
		return nil, fmt.Errorf("[%v]=[%d is invalid_year]", createBookDir, book.Year)
	}

	err = pgTx.QueryRow(
		ctx,
		createBookQuery,
		book.Title,
		book.Author_ID,
		book.Year,
		book.Genre).Scan(&book.ID)
	if err != nil {
		return nil, fmt.Errorf("[%v.QueryRow2]=[%v]", createBookDir, err)
	}

	err = pgTx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%v.Commit]=[%v]", createBookDir, err)
	}

	return models.ConvertBookDAOToDTO(book), nil
}

// GetBookBYID is
func (r *BookRepository) GetBookByID(ctx context.Context, book_id int) (*models.BookDTO, error) {
	var Book models.BookDAO
	err := pgxscan.Get(
		ctx,
		r.DB,
		&Book,
		gettingOneBookQuery,
		book_id)
	if err != nil {
		return nil, fmt.Errorf("[%v.Get]=[%v]", getBookByIDDir, err)
	}

	return models.ConvertBookDAOToDTO(&Book), nil
}

// Author can be search books about with title or published year
// Paginnation also need; generally filter need::
// GetAllBooks is
func (r *BookRepository) GetAllBooks(ctx context.Context, pagination models.PaginationForBook) ([]*models.BookDTO, error) {
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

	err := pgxscan.Select(
		ctx,
		r.DB,
		&Books,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("[%v.Select]=[%v]", getAllBooksDir, err)
	}

	var BookDTOs []*models.BookDTO
	for _, book := range Books {
		BookDTOs = append(BookDTOs, models.ConvertBookDAOToDTO(book))
	}

	return BookDTOs, nil
}

// UpdateBook is
func (r *BookRepository) UpdateBook(ctx context.Context, book_id int, updateInput *models.UpdateInputBook) (string, error) {
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
		return "", fmt.Errorf("[%v.checkingLenSetValues]", updateBookDir)
	}

	query := fmt.Sprintf(`UPDATE books SET
								%s WHERE id = $%d
								RETURNING 'Book informations updated'
								`, strings.Join(setValues, ", "), argId)
	args = append(args, book_id)

	var response string
	_, err := r.DB.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return response, fmt.Errorf("[%v.Exec]=[%v]", updateBookDir, err)
	}

	response = fmt.Sprintf("Book with ID %d updated successfully", book_id)
	return response, nil
}

// DeleteBook is
func (r *BookRepository) DeleteBook(ctx context.Context, book_id int) (string, error) {
	var response string
	err := r.DB.QueryRow(
		ctx,
		deleteBookQuery,
		book_id).Scan(&response)
	if err != nil {
		return response, fmt.Errorf("[%v.QueryRow]=[%v]", deleteBookDir, err)
	}

	response = fmt.Sprintf("Book with ID %d deleted successfully", book_id)
	return response, nil
}
