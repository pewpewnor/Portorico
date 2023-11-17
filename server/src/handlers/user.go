package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/src/model/response"
)

func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userRepository.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.InternalProblem("server cannot get all users"))
	}

	return c.Status(http.StatusOK).JSON(response.Success("successfully found all users", users))
}

func (h *handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	if ok, res := h.BodyParseAndValidate(c, &body); !ok {
		return res
	}

	user, _, err := h.userRepository.Create(body.Username, body.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.InternalProblem("server cannot create user"))
	}

	return c.Status(http.StatusOK).JSON(response.Success("successfully created user", user))
}
