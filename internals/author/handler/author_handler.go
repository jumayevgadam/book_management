package handler

import (
	"strconv"
	"time"

	"github.com/jumayevgadam/book_management/internals/author/models"
	"github.com/jumayevgadam/book_management/internals/author/service"
	httperr "github.com/jumayevgadam/book_management/pkg/httpErr"
	"github.com/jumayevgadam/book_management/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type AuthorHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "author.handler.CreateAuthor")
		defer span.Finish()

		var Author models.AuthorDAO
		if err := c.Bind(&Author); err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		data, err := h.service.CreateAuthor(ctx, &Author)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil
	}
}

// GetAuthorByID is
func (h *AuthorHandler) GetAuthorByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "author.handler.GetAuthorByID")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		data, err := h.service.GetAuthorByID(ctx, id)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil
	}
}

// GetAllAuthors is
func (h *AuthorHandler) GetAllAuthors() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "author.handler.GetAllAuthor")
		defer span.Finish()

		var pagination models.PaginationForAuthor
		limit := c.Request().URL.Query().Get("limit")
		if limit == "" {
			limit = "10" // default value
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt <= 0 {
			return c.JSON(httperr.ErrorResponse(err))
		}
		pagination.Limit = limitInt

		offset := c.Request().URL.Query().Get("offset")
		if offset == "" {
			offset = "0" // default value
		}

		offsetInt, err := strconv.Atoi(offset)
		if err != nil || offsetInt < 0 {
			return c.JSON(httperr.ErrorResponse(err))
		}
		pagination.Offset = offsetInt

		criteria := c.Request().URL.Query().Get("criteria")
		pagination.Criteria = criteria

		authors, err := h.service.GetAllAuthor(ctx, pagination)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"data": authors,
		})

		return nil
	}
}

// UpdateAuthor is
func (h *AuthorHandler) UpdateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "author.handler.UpdateAuthor")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
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
				return c.JSON(httperr.ErrorResponse(err))
			}
			updateInput.Birthdate = &parsedDate
		}

		responseData, err := h.service.UpdateAuthor(ctx, id, &updateInput)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"response": responseData,
		})

		return nil
	}
}

// DeleteAuthor is
func (h *AuthorHandler) DeleteAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "author.handler.DeleteAuthor")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		responseData, err := h.service.DeleteAuthor(ctx, id)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"responseData": responseData,
		})

		return nil
	}
}
