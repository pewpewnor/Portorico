package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) ServerStatus(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]any{"is_healthy": true})
}
