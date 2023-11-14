package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/model/response"
)

func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepository.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.InternalProblem("server cannot get all users"))
	}

	return c.Status(http.StatusOK).JSON(response.Success("successfully found all users", users))
}

func (h *handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
	}
	if res := h.BodyParseAndValidate(c, &body); res != nil {
		return res
	}

	user, err := h.UserRepository.Create(body.Username, body.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.InternalProblem("server cannot create user"))
	}

	return c.Status(http.StatusOK).JSON(response.Success("successfully created user", user))
}
