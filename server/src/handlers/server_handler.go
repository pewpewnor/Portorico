package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) ServerStatus(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(map[string]any{"is_healthy": true})
}
