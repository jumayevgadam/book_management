package models

func ConvertAuthorDAOToDTO(author *AuthorDAO) *AuthorDTO {
	return &AuthorDTO{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		Birthdate: author.Birthdate,
	}
}

func ConvertAuthorDTOToDAO(author *AuthorDTO) *AuthorDAO {
	return &AuthorDAO{
		ID:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
		Birthdate: author.Birthdate,
	}
}
