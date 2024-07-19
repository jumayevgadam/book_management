package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/service"
	response "github.com/jumayevgadam/book_management/pkg/customerr"
)

type AuthorHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var Author models.Author

	name := c.PostForm("name")
	if name == "" {
		response.NewError(c, 400, "Name is required")
		return
	}
	Author.Name = name

	biography := c.PostForm("biography")
	Author.Biography = biography

	birthdatestr := c.PostForm("birthdate")
	birthdate, err := time.Parse("2006-01-02", birthdatestr)
	if err != nil {
		response.NewError(c, 400, "Invalid birthday format")
		return
	}
	Author.Birthdate = birthdate

	data, err := h.service.CreateAuthor(c, &Author)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, 400, "Invalid author id")
		return
	}

	data, err := h.service.GetAuthorByID(c, id)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *AuthorHandler) GetAllAuthors(c *gin.Context) {
	var pagination models.PaginationForAuthor

	limit := c.Request.URL.Query().Get("limit")
	if limit == "" {
		limit = "10" // default value
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		response.NewError(c, 400, err.Error())
		return
	}
	pagination.Limit = limitInt

	offset := c.Request.URL.Query().Get("offset")
	if offset == "" {
		offset = "0" // default value
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		response.NewError(c, 400, err.Error())
		return
	}
	pagination.Offset = offsetInt

	criteria := c.Request.URL.Query().Get("criteria")
	pagination.Criteria = criteria

	authors, err := h.service.GetAllAuthor(c, pagination)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": authors,
	})
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, 400, "Invalid author id")
		return
	}

	var updateInput models.UpdateInputAuthor

	if name := c.PostForm("name"); name != "" {
		updateInput.Name = &name
	}

	if biography := c.PostForm("biography"); biography != "" {
		updateInput.Biography = &biography
	}

	if birthdateStr := c.PostForm("birthdate"); birthdateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", birthdateStr)
		if err != nil {
			response.NewError(c, 400, "error in parsing birthdate")
			return
		}
		updateInput.Birthdate = &parsedDate
	}

	responseData, err := h.service.UpdateAuthor(c, id, &updateInput)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"response": responseData,
	})
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid author ID"})
		return
	}

	responseData, err := h.service.DeleteAuthor(c, id)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"response": responseData,
	})
}
