package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/book/service"
)

type BookHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *BookHandler) GetBookByID() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *BookHandler) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// UpdateBook is
func (h *BookHandler) UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// DeleteBook is
func (h *BookHandler) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
