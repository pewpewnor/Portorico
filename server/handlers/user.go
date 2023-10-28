package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pewpewnor/portorico/server/model"
	"github.com/pewpewnor/portorico/server/response"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (h *Handler) GetAllUsers(c echo.Context) error {
	var users []model.User
	if err := h.DB.Find(&users).Error; err != nil {
		log.Printf("Server cannot get all users: %v\n", err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if ok := h.BindAndValidate(c, &req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user := &model.User{Name: req.Name}
	if err := h.DB.Create(&user).Error; err != nil {
		log.Printf("Server cannot create user: %v\n", err)
		return c.JSON(http.StatusInternalServerError, response.SError("Server cannot create user"))
	}

	return c.JSON(http.StatusCreated, response.Success("Successfully created user", user))
}
