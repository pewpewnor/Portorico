package handlers

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/model"
	"github.com/pewpewnor/portorico/server/response"
	"github.com/pewpewnor/portorico/server/validator"
)

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	if err := h.DB.Find(&users).Error; err != nil {
		log.Errorf("Server cannot get all users: %v\n", err)
	}

	return c.Status(http.StatusOK).JSON(users)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.SErrorFromErr("Request malformed", err))
	}
	if validations := validator.Validate(&req); len(validations) > 0 {
		return c.Status(http.StatusBadRequest).JSON(response.Error("Request malformed", "Validation failed", validations))
	}

	user := &model.User{Name: req.Name}
	if err := h.DB.Create(&user).Error; err != nil {
		log.Warnf("Server cannot create user: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(response.SError("Server cannot create user"))
	}

	return c.Status(http.StatusOK).JSON(response.Success("Successfully created user", user))
}
