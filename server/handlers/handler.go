package handlers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/model/response"
	"github.com/pewpewnor/portorico/server/repository"
	"gorm.io/gorm"
)

type handler struct {
	DB             *gorm.DB
	Validator      *validator.Validate
	UserRepository *repository.UserRepository
}

func NewHandler(db *gorm.DB, validator *validator.Validate) *handler {
	return &handler{
		DB:             db,
		Validator:      validator,
		UserRepository: &repository.UserRepository{DB: db},
	}
}

func (h *handler) validate(data any) []response.FieldValidation {
	validations := []response.FieldValidation{}

	errs := h.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validations = append(validations, response.FieldValidation{
				Field:         err.Field(),
				ReceivedValue: err.Value(),
				Message:       "Validation failed for: " + err.Tag(),
			})
		}
	}

	return validations
}

func (h *handler) BodyParseAndValidate(c *fiber.Ctx, dataPtr any) error {
	if err := c.BodyParser(dataPtr); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.RequestMalformed("request body is malformed"))
	}
	if validations := h.validate(dataPtr); len(validations) > 0 {
		return c.Status(http.StatusBadRequest).JSON(response.RequestMalformedWithValidations(validations))
	}

	return nil
}
