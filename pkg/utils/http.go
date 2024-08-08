package utils

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

// Get Request ID from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

type ReqIDCtxKey struct{}

// Get ctx with timeout and request id from the echo context
func GetCtxWithReqID(c echo.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// Get Context with Request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// Get Config Path
func GetConfigPath(configPath string) string {
	if configPath == "./config/config-docker" {
		return "./config/config-docker"
	}

	return "./config/config-local"
}

// Configure JWT cookie
