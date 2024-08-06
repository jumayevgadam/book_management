package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/handler"
	"github.com/jumayevgadam/book_management/internals/book/repository"
	"github.com/jumayevgadam/book_management/internals/book/service"
	"github.com/jumayevgadam/book_management/pkg/logger"
	"github.com/labstack/echo/v4"
)

func InitBookRoutes(router *echo.Group, DB *pgxpool.Pool, logger logger.Logger) {
	// Book routes
	BookRepos := repository.NewDTORepository(DB, logger)
	BookServices := service.NewDTOService(BookRepos)
	BookHandlers := handler.NewDTOHandler(BookServices)

	BookRoutes := router.Group("/book")
	{
		BookRoutes.POST("", BookHandlers.CreateBook())
		BookRoutes.GET("", BookHandlers.GetAllBooks())
		BookRoutes.GET("/:id", BookHandlers.GetBookByID())
		BookRoutes.PUT("/:id", BookHandlers.UpdateBook())
		BookRoutes.DELETE("/:id", BookHandlers.DeleteBook())
	}
}
