package handlers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/response"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
	Validator *validator.Validate
}

func (h *Handler) validate(data any) []response.FieldValidation {
	validations := []response.FieldValidation{}

	errs := h.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validations = append(validations, response.FieldValidation{
				Field:         err.Field(),
				ReceivedValue: err.Value(),
				Message:       "Field is supposed to be " + err.Tag(),
			})
		}
	}

	return validations
}

func (h *Handler) BodyParseAndValidate(c *fiber.Ctx, dataPtr any) (bool, error) {
	if err := c.BodyParser(dataPtr); err != nil {
		return false, c.Status(http.StatusBadRequest).JSON(response.SErrorFromErr("Request malformed", err))
	}
	if validations := h.validate(dataPtr); len(validations) > 0 {
		return false, c.Status(http.StatusBadRequest).JSON(response.Error("Request malformed", "Validation failed", validations))
	}
	return true, nil
}