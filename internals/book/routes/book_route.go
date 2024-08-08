package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/handler"
	"github.com/jumayevgadam/book_management/internals/book/repository"
	"github.com/jumayevgadam/book_management/internals/book/service"
	"github.com/labstack/echo/v4"
)

func InitBookRoutes(router *echo.Group, DB *pgxpool.Pool) {
	// Data Flows from repo -> service -> handler
	BookRepos := repository.NewDTORepository(DB)
	BookServices := service.NewDTOService(BookRepos)
	BookHandlers := handler.NewDTOHandler(BookServices)
	// -> routes
	// BookRoutes are
	BookRoutes := router.Group("/book")
	{
		BookRoutes.POST("", BookHandlers.CreateBook())
		BookRoutes.GET("", BookHandlers.GetAllBooks())
		BookRoutes.GET("/:id", BookHandlers.GetBookByID())
		BookRoutes.PUT("/:id", BookHandlers.UpdateBook())
		BookRoutes.DELETE("/:id", BookHandlers.DeleteBook())
	}
}
