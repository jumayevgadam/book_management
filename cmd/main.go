package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/book_management/internals/dbconn"

	"github.com/joho/godotenv"
	"github.com/jumayevgadam/book_management/internals/author/routes"
	routes2 "github.com/jumayevgadam/book_management/internals/book/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables from.env file
	if e := godotenv.Load(); e != nil {
		logrus.Printf("failed to initialize .env file: %v", e)
	}

	// GettingDBClient
	DB, err := dbconn.GetDBClient()
	if err != nil {
		logrus.Printf("failed to connect to database: %v", err)
	}

	// New Echo
	app := gin.Default()
	api := app.Group("/api")

	// Initialize routes
	routes.InitAuthorRoutes(api, DB)
	routes2.InitBookRoutes(api, DB)

	// Call server
	app.Run(":3000")
}
