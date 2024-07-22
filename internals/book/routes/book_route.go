package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/handler"
	"github.com/jumayevgadam/book_management/internals/book/repository"
	"github.com/jumayevgadam/book_management/internals/book/service"
	"github.com/jumayevgadam/book_management/pkg/logger"
)

func InitBookRoutes(router *gin.RouterGroup, DB *pgxpool.Pool, logger logger.Logger) {
	// Book routes
	BookRepos := repository.NewDTORepository(DB, logger)
	BookServices := service.NewDTOService(BookRepos)
	BookHandlers := handler.NewDTOHandler(BookServices)

	BookRoutes := router.Group("/books")
	{
		BookRoutes.POST("", BookHandlers.CreateBook)
		BookRoutes.GET("", BookHandlers.GetAllBooks)
		BookRoutes.GET("/:id", BookHandlers.GetBookByID)
		BookRoutes.PUT("/:id", BookHandlers.UpdateBook)
		BookRoutes.DELETE("/:id", BookHandlers.DeleteBook)
	}
}
