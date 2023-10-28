package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ServerStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"is_healthy": true})
}
