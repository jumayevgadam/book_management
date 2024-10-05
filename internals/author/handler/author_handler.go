package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/service"
)

type AuthorHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) CreateAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Author models.AuthorDAO
		if err := ctx.ShouldBind(&Author); err != nil {
			ctx.JSON(400, err.Error())
			// data, err := h.service.CreateAuthor(ctx, &Author)
			// if err != nil {
			// 	return c.JSON(httperr.ErrorResponse(err))
			// }

			// c.JSON(200, echo.Map{
			// 	"data": data,
			// })
		}
	}
}

// GetAuthorByID is
func (h *AuthorHandler) GetAuthorByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
	}
}

// GetAllAuthors is
func (h *AuthorHandler) GetAllAuthors() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// UpdateAuthor is
func (h *AuthorHandler) UpdateAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// DeleteAuthor is
func (h *AuthorHandler) DeleteAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
