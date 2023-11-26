package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *handler) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("session")
	if token == "" {
		return c.SendStatus(401)
	}

	user, found := h.userRepo.GetBySessionToken(token)
	if !found {
		return c.SendStatus(401)
	}

	c.Locals("user", user)
	return c.Next()
}
