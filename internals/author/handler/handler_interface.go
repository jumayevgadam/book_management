package handler

import "github.com/labstack/echo/v4"

type IHandler interface {
	CreateAuthor() echo.HandlerFunc
	GetAuthorByID() echo.HandlerFunc
	GetAllAuthors() echo.HandlerFunc
	UpdateAuthor() echo.HandlerFunc
	DeleteAuthor() echo.HandlerFunc
}
