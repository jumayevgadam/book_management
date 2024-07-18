package handler

import "github.com/jumayevgadam/book_management/internals/author/service"

type AuthorHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *AuthorHandler {
	return &AuthorHandler{service: service}
}
