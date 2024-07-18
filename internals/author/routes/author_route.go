package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book_management/internals/author/handler"
	"github.com/jumayevgadam/book_management/internals/author/repository"
	"github.com/jumayevgadam/book_management/internals/author/service"
)

func InitAuthorRoutes(router *gin.RouterGroup, DB *pgxpool.Pool) {
	AuthorRepos := repository.NewDTORepository(DB)
	AuthorServices := service.NewDTOService(AuthorRepos)
	AuthorHandlers := handler.NewDTOHandler(AuthorServices)

	AuthorRoutes := router.Group("/author")
	{
		AuthorRoutes.POST("/create", AuthorHandlers.CreateAuthor)
		AuthorRoutes.GET("/get/:id", AuthorHandlers.GetAuthorByID)
		AuthorRoutes.GET("/get/all", AuthorHandlers.GetAllAuthors)
		AuthorRoutes.PUT("/update/:id", AuthorHandlers.UpdateAuthor)
		AuthorRoutes.DELETE("/delete/:id", AuthorHandlers.DeleteAuthor)
	}
}
