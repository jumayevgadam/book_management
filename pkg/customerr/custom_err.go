package customerr

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewError(c echo.Context, statusCode int, message string) {
	logrus.Error(message)
	c.JSON(statusCode, errorResponse{message})
	//c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
