package models

func ConvertBookDAOToDTO(book *BookDAO) *BookDTO {
	return &BookDTO{
		ID:        book.ID,
		Title:     book.Title,
		Author_ID: book.Author_ID,
		Year:      book.Year,
		Genre:     book.Genre,
	}
}

func ConvertBookDTOToDAO(book *BookDTO) *BookDAO {
	return &BookDAO{
		ID:        book.ID,
		Title:     book.Title,
		Author_ID: book.Author_ID,
		Year:      book.Year,
		Genre:     book.Genre,
	}
}
