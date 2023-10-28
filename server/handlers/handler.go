package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pewpewnor/portorico/server/response"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) BindAndValidate(c echo.Context, target any) bool {
	if err := c.Bind(target); err != nil {
		c.JSON(http.StatusBadRequest, response.SErrorFromErr("Request malformed", err))
		return false
	}

	if err := c.Validate(target); err != nil {
		validations := []response.ErrorResponseValidation{}
		for _, line := range strings.Split(err.Error(), "\n") {
			vals := strings.Split(line, ":")
			fieldVals := strings.Split(vals[1], ".")[1]
			validations = append(validations, response.NewValidation(fieldVals[:len(fieldVals)-7], vals[2]))
		}

		c.JSON(http.StatusBadRequest, response.Error("Request malformed", "Validation failed", validations))
		return false
	}

	return true
}
