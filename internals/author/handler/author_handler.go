package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/service"
	"github.com/sirupsen/logrus"
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
		c.JSON(400, gin.H{"error": "Name is required"})
		return
	}
	Author.Name = name

	biography := c.PostForm("biography")
	Author.Biography = biography

	birthdatestr := c.PostForm("birthdate")
	birthdate, err := time.Parse("2006-01-02", birthdatestr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid birthday format"})
		return
	}
	Author.Birthdate = birthdate

	data, err := h.service.CreateAuthor(c, &Author)
	if err != nil {
		logrus.Errorf("error occured in creating author: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid author ID"})
		return
	}

	data, err := h.service.GetAuthorByID(c, id)
	if err != nil {
		logrus.Errorf("error occured in fetching author: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *AuthorHandler) GetAllAuthors(c *gin.Context) {
	authors, err := h.service.GetAllAuthor(c)
	if err != nil {
		logrus.Errorf("error occured in fetching all authors: %v", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": authors,
	})
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid author ID"})
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
			logrus.Errorf("error in parsing birthdate: %v", err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		updateInput.Birthdate = &parsedDate
	}

	response, err := h.service.UpdateAuthor(c, id, &updateInput)
	if err != nil {
		logrus.Errorf("error occured in updating author: %v", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": response,
	})
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid author ID"})
		return
	}

	response, err := h.service.DeleteAuthor(c, id)
	if err != nil {
		logrus.Errorf("error occured in deleting author: %v", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": response,
	})
}
