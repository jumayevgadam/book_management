package main

import (
	"os"

	"github.com/jumayevgadam/book_management/internals/dbconn"
	"github.com/jumayevgadam/book_management/internals/server"
	"github.com/jumayevgadam/book_management/pkg/logger"
	"github.com/labstack/echo/v4"

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

	DB, err := dbconn.GetDBClient()
	if err != nil {
		logrus.Printf("failed to connect to database: %v", err)
	}

	log := logger.NewLogrusLogger()

	app := echo.New()
	api := app.Group("/api")

	// Initialize routes
	routes.InitAuthorRoutes(api, DB, log)
	routes2.InitBookRoutes(api, DB, log)

	// Call server
	srv := &server.Server{}
	if err := srv.Run(os.Getenv("PORT_SERVER"), app); err != nil {
		logrus.Printf("failed to run server: %v", err.Error())
	}
}
