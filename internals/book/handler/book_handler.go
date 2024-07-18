package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/service"
	"github.com/sirupsen/logrus"
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
		c.JSON(400, gin.H{"error": "Title is required"})
		return
	}
	Book.Title = title

	authorID, err := strconv.Atoi(c.PostForm("author_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid author ID"})
		return
	}
	Book.Author_ID = authorID

	year, err := strconv.Atoi(c.PostForm("year"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid year"})
		return
	}
	Book.Year = year

	genre := c.PostForm("genre")
	Book.Genre = genre

	data, err := h.service.CreateBook(c, &Book)
	if err != nil {
		logrus.Errorf("failed to create book: %v", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid author id"})
		return
	}

	data, err := h.service.GetBookByID(c, id)
	if err != nil {
		logrus.Errorf("failed to get book: %v", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {

}

func (h *BookHandler) UpdateBook(c *gin.Context) {

}

func (h *BookHandler) DeleteBook(c *gin.Context) {

}
