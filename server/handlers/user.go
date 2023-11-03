package handlers

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/model/response"
)

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepository.GetAll()
	if err != nil {
		log.Errorf("Server cannot get all users: %v\n", err)
	}

	return c.Status(http.StatusOK).JSON(response.Success("successfully found all users", users))
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username" validate:"bodyuired"`
		Password string `json:"password" validate:"bodyuired"`
		Name     string `json:"name" validate:"bodyuired"`
	}
	if res := h.BodyParseAndValidate(c, &body); res != nil {
		return res
	}

	user, err := h.UserRepository.Create(body.Username, body.Password, body.Name)
	if err != nil {
		log.Warnf("Server cannot create user: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(response.InternalProblem("server cannot create user"))
	}

	return c.Status(http.StatusOK).JSON(response.Success("Successfully created user", user))
}
