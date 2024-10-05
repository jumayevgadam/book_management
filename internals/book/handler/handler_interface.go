package handler

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	CreateBook() gin.HandlerFunc
	GetBookByID() gin.HandlerFunc
	GetAllBooks() gin.HandlerFunc
	UpdateBook() gin.HandlerFunc
	DeleteBook() gin.HandlerFunc
}
