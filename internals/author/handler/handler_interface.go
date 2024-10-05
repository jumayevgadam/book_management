package handler

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	CreateAuthor() gin.HandlerFunc
	GetAuthorByID() gin.HandlerFunc
	GetAllAuthors() gin.HandlerFunc
	UpdateAuthor() gin.HandlerFunc
	DeleteAuthor() gin.HandlerFunc
}
