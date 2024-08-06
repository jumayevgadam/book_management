package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/handler"
	"github.com/jumayevgadam/book_management/internals/author/repository"
	"github.com/jumayevgadam/book_management/internals/author/service"
	"github.com/jumayevgadam/book_management/pkg/logger"
	"github.com/labstack/echo/v4"
)

func InitAuthorRoutes(router *echo.Group, DB *pgxpool.Pool, logger logger.Logger) {
	// These codes implementations of code such a river))
	AuthorRepos := repository.NewDTORepository(DB, logger)
	AuthorServices := service.NewDTOService(AuthorRepos)
	AuthorHandlers := handler.NewDTOHandler(AuthorServices)

	AuthorRoutes := router.Group("/authors")
	{
		AuthorRoutes.POST("", AuthorHandlers.CreateAuthor())
		AuthorRoutes.GET("", AuthorHandlers.GetAllAuthors())
		AuthorRoutes.GET("/:id", AuthorHandlers.GetAuthorByID())
		AuthorRoutes.PUT("/:id", AuthorHandlers.UpdateAuthor())
		AuthorRoutes.DELETE("/:id", AuthorHandlers.DeleteAuthor())
	}
}
