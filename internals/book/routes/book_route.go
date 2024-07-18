package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/book/handler"
	"github.com/jumayevgadam/book_management/internals/book/repository"
	"github.com/jumayevgadam/book_management/internals/book/service"
)

func InitBookRoutes(router *gin.RouterGroup, DB *pgxpool.Pool) {
	// Book routes
	BookRepos := repository.NewDTORepository(DB)
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
