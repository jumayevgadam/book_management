package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/service"
	response "github.com/jumayevgadam/book_management/pkg/customerr"
)

type BookHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var Book models.Book

	title := c.PostForm("title")
	if title == "" {
		response.NewError(c, 400, "title is required")
		return
	}
	Book.Title = title

	authorID, err := strconv.Atoi(c.PostForm("author_id"))
	if err != nil {
		response.NewError(c, 400, "invalid author id")
		return
	}
	Book.Author_ID = authorID

	year, err := strconv.Atoi(c.PostForm("year"))
	if err != nil {
		response.NewError(c, 400, err.Error())
		return
	}
	Book.Year = year

	genre := c.PostForm("genre")
	Book.Genre = genre

	data, err := h.service.CreateBook(c, &Book)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, 400, "invalid author id")
		return
	}

	data, err := h.service.GetBookByID(c, id)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	var pagination models.PaginationForBook

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
	if err != nil {
		response.NewError(c, 400, err.Error())
		return
	}
	pagination.Offset = offsetInt

	title := c.Request.URL.Query().Get("title")
	pagination.Title = title

	yearStr := c.Request.URL.Query().Get("year")
	yearInt, err := strconv.Atoi(yearStr)
	if err != nil && yearInt < 0 && yearInt > time.Now().Year() {
		response.NewError(c, 400, "invalid year")
		return
	}
	pagination.Year = yearInt

	genre := c.Request.URL.Query().Get("genre")
	pagination.Genre = genre

	data, err := h.service.GetAllBooks(c, pagination)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})

}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, 400, "Invalid book id")
		return
	}

	var updateInput models.UpdateInputBook

	if title := c.PostForm("title"); title != "" {
		updateInput.Title = &title
	}

	if year := c.PostForm("year"); year != "" {
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			response.NewError(c, 400, err.Error())
			return
		}

		if yearInt > time.Now().Year() {
			response.NewError(c, 400, "invalid year")
			return
		}

		updateInput.Year = &yearInt
	}

	if genre := c.PostForm("genre"); genre != "" {
		updateInput.Genre = &genre
	}

	responseData, err := h.service.UpdateBook(c, id, &updateInput)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"response": responseData,
	})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(c, 400, "Invalid book id")
		return
	}

	responseData, err := h.service.DeleteBook(c, id)
	if err != nil {
		response.NewError(c, 500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"response": responseData,
	})
}
