package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userRepository.GetAll()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(map[string]any{"users": users})
}

func (h *handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	validations := map[string]string{}
	if body.Username == "" {
		validations["username"] = "username must not be empty"
	}
	if body.Password == "" {
		validations["password"] = "password must not be empty"
	}
	if len(validations) == 0 {
		return c.Status(http.StatusBadRequest).JSON(validations)
	}

	user := h.userRepository.GetByUsername(body.Username)
	if user != nil {
		validations["username"] =
			"username is already taken, please try another one"
		return c.Status(http.StatusBadRequest).JSON(
			map[string]any{"validations": validations})
	}

	user, _, err := h.userRepository.Create(body.Username, body.Password)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(map[string]any{"user": user})
}
