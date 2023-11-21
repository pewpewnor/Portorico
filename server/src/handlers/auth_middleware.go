package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.SendStatus(400)
	}

	user := h.userRepository.GetBySessionToken(token)
	if user == nil {
		return c.SendStatus(401)
	}

	c.Locals("user", user)
	return c.Next()
}
