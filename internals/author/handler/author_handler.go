package handler

import (
	"strconv"
	"time"

	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/service"
	response "github.com/jumayevgadam/book_management/pkg/customerr"
	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Author models.AuthorDAO

		name := c.FormValue("name")
		if name == "" {
			response.NewError(c, 400, "Name is required")
			return nil
		}
		Author.Name = name

		biography := c.FormValue("biography")
		Author.Biography = biography

		birthdatestr := c.FormValue("birthdate")
		birthdate, err := time.Parse("2006-01-02", birthdatestr)
		if err != nil {
			response.NewError(c, 400, "Invalid birthday format")
			return err
		}
		Author.Birthdate = birthdate

		data, err := h.service.CreateAuthor(c, &Author)
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

func (h *AuthorHandler) GetAuthorByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewError(c, 400, "Invalid author id")
			return err
		}

		data, err := h.service.GetAuthorByID(c, id)
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

func (h *AuthorHandler) GetAllAuthors() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pagination models.PaginationForAuthor

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
		if err != nil || offsetInt < 0 {
			response.NewError(c, 400, err.Error())
			return err
		}
		pagination.Offset = offsetInt

		criteria := c.Request().URL.Query().Get("criteria")
		pagination.Criteria = criteria

		authors, err := h.service.GetAllAuthor(c, pagination)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"data": authors,
		})

		return nil
	}
}

func (h *AuthorHandler) UpdateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewError(c, 400, "Invalid author id")
			return err
		}

		var updateInput models.UpdateInputAuthor

		if name := c.FormValue("name"); name != "" {
			updateInput.Name = &name
		}

		if biography := c.FormValue("biography"); biography != "" {
			updateInput.Biography = &biography
		}

		if birthdateStr := c.FormValue("birthdate"); birthdateStr != "" {
			parsedDate, err := time.Parse("2006-01-02", birthdateStr)
			if err != nil {
				response.NewError(c, 400, "error in parsing birthdate")
				return err
			}
			updateInput.Birthdate = &parsedDate
		}

		responseData, err := h.service.UpdateAuthor(c, id, &updateInput)
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

func (h *AuthorHandler) DeleteAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.NewError(c, 400, "error converting to integer")
			return err
		}

		responseData, err := h.service.DeleteAuthor(c, id)
		if err != nil {
			response.NewError(c, 500, err.Error())
			return err
		}

		c.JSON(200, echo.Map{
			"responseData": responseData,
		})

		return nil
	}
}
