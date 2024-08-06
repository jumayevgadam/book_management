package handler

import "github.com/labstack/echo/v4"

type IHandler interface {
	CreateBook() echo.HandlerFunc
	GetBookByID() echo.HandlerFunc
	GetAllBooks() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}
