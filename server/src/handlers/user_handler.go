package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) Register(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	validations := map[string]string{}
	h.validateStringMaxLength(validations, "username", "username", 64, body.Username)
	h.validateStringMinMaxLength(validations, "password", "password", 6, 64, body.Password)
	if len(validations) > 0 {
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	_, exist := h.userRepo.GetByUsername(body.Username)
	if exist {
		validations["username"] = "username is already taken, please try a different one"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	user, session, err := h.userRepo.Create(body.Username, body.Password)
	if err != nil {
		return c.SendStatus(500)
	}

	c.Cookie(&fiber.Cookie{Name: "session", Value: session.Token, Expires: time.Now().Add(24 * time.Hour)})
	return c.Status(200).JSON(map[string]any{"user": user})
}

func (h *handler) Login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	validations := map[string]string{}
	h.validateStringMaxLength(validations, "username", "username", 64, body.Username)
	h.validateStringMaxLength(validations, "password", "password", 64, body.Password)
	if len(validations) > 0 {
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	user, session, valid, err := h.userRepo.GetByCredentials(body.Username, body.Password)
	if err != nil {
		return c.SendStatus(500)
	}
	if !valid {
		validations["general"] = "username or password is incorrect, please try again"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	c.Cookie(&fiber.Cookie{Name: "session", Value: session.Token, Expires: time.Now().Add(24 * time.Hour)})
	return c.Status(200).JSON(map[string]any{"user": user})
}
