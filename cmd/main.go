package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jumayevgadam/book_management/internals/author/routes"
	"github.com/jumayevgadam/book_management/pkg/dbconn"
	"github.com/sirupsen/logrus"
)

func main() {
	if e := godotenv.Load(); e != nil {
		logrus.Fatalf("failed to initialize .env file: %v", e)
	}

	DB, err := dbconn.GetDBClient()
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
	}

	app := gin.Default()

	api := app.Group("/api")

	routes.InitAuthorRoutes(api, DB)

}
