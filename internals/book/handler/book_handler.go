package handler

import (
	"strconv"
	"time"

	"github.com/jumayevgadam/book_management/internals/book/models"
	"github.com/jumayevgadam/book_management/internals/book/service"
	httperr "github.com/jumayevgadam/book_management/pkg/httpErr"
	"github.com/jumayevgadam/book_management/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type BookHandler struct {
	service *service.Service
}

func NewDTOHandler(service *service.Service) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "book.handler.CreateBook")
		defer span.Finish()

		var Book models.BookDAO
		if err := c.Bind(&Book); err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		data, err := h.service.CreateBook(ctx, &Book)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil
	}
}

func (h *BookHandler) GetBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "book.handler.GetBookByID")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		data, err := h.service.GetBookByID(ctx, id)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil
	}
}

func (h *BookHandler) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "book.handler.GetAllBooks")
		defer span.Finish()

		var pagination models.PaginationForBook
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
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}
		pagination.Offset = offsetInt

		title := c.Request().URL.Query().Get("title")
		pagination.Title = title

		yearStr := c.Request().URL.Query().Get("year")
		yearInt, err := strconv.Atoi(yearStr)
		if err != nil && yearInt < 0 && yearInt > time.Now().Year() {
			return c.JSON(httperr.ErrorResponse(err))
		}
		pagination.Year = yearInt

		genre := c.Request().URL.Query().Get("genre")
		pagination.Genre = genre

		data, err := h.service.GetAllBooks(ctx, pagination)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"data": data,
		})

		return nil

	}
}

// UpdateBook is
func (h *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "book.handler.UpdateBook")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		var updateInput models.UpdateInputBook
		if err := c.Bind(&updateInput); err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		responseData, err := h.service.UpdateBook(ctx, id, &updateInput)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"response": responseData,
		})

		return nil
	}
}

// DeleteBook is
func (h *BookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "book.handler.DeleteBook")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		responseData, err := h.service.DeleteBook(ctx, id)
		if err != nil {
			return c.JSON(httperr.ErrorResponse(err))
		}

		c.JSON(200, echo.Map{
			"response": responseData,
		})

		return nil
	}
}
