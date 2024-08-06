package handler

import (
	"strconv"
	"time"

	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/service"
	response "github.com/jumayevgadam/book_management/pkg/customerr"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Book models.BookDAO

		title := c.FormValue("title")
		if title == "" {
			response.NewError(c, 400, "title is required")
			return nil
		}
		Book.Title = title

		authorID, err := strconv.Atoi(c.FormValue("author_id"))
		if err != nil {
			response.NewError(c, 400, "invalid author id")
			return err
		}
		Book.Author_ID = authorID

		year, err := strconv.Atoi(c.FormValue("year"))
		if err != nil {
			response.NewError(c, 400, err.Error())
			return err
		}
		Book.Year = year

		genre := c.FormValue("genre")
		Book.Genre = genre

		data, err := h.service.CreateBook(c, &Book)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil
	}
}

func (h *BookHandler) GetBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewError(c, 400, "invalid author id")
			return nil
		}

		data, err := h.service.GetBookByID(c, id)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil
	}
}

func (h *BookHandler) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pagination models.PaginationForBook

		limit := c.Request().URL.Query().Get("limit")
		if limit == "" {
			limit = "10" // default value
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt <= 0 {
			response.NewError(c, 400, err.Error())
			return err
		}
		pagination.Limit = limitInt

		offset := c.Request().URL.Query().Get("offset")
		if offset == "" {
			offset = "0" // default value
		}

		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			response.NewError(c, 400, err.Error())
			return err
		}
		pagination.Offset = offsetInt

		title := c.Request().URL.Query().Get("title")
		pagination.Title = title

		yearStr := c.Request().URL.Query().Get("year")
		yearInt, err := strconv.Atoi(yearStr)
		if err != nil && yearInt < 0 && yearInt > time.Now().Year() {
			response.NewError(c, 400, "invalid year")
			return err
		}
		pagination.Year = yearInt

		genre := c.Request().URL.Query().Get("genre")
		pagination.Genre = genre

		data, err := h.service.GetAllBooks(c, pagination)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil

	}
}

func (h *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewError(c, 400, "Invalid book id")
			return err
		}

		var updateInput models.UpdateInputBook

		if title := c.FormValue("title"); title != "" {
			updateInput.Title = &title
		}

		if year := c.FormValue("year"); year != "" {
			yearInt, err := strconv.Atoi(year)
			if err != nil {
				response.NewError(c, 400, err.Error())
				return err
			}

			if yearInt > time.Now().Year() {
				response.NewError(c, 400, "invalid year")
				return err
			}

			updateInput.Year = &yearInt
		}

		if genre := c.FormValue("genre"); genre != "" {
			updateInput.Genre = &genre
		}

		responseData, err := h.service.UpdateBook(c, id, &updateInput)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"response": responseData,
		})

		return nil
	}
}

func (h *BookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewError(c, 400, "Invalid book id")
			return err
		}

		responseData, err := h.service.DeleteBook(c, id)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"response": responseData,
		})

		return nil
	}
}
