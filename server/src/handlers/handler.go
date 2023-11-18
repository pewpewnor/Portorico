package handlers

import (
	"net/http"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pewpewnor/portorico/server/src/model/response"
	"github.com/pewpewnor/portorico/server/src/repository"
)

type handler struct {
	userRepository *repository.LiveUserRepository
	validator      *validator.Validate
}

func NewHandler(db *sqlx.DB, validator *validator.Validate) *handler {
	return &handler{repository.NewLiveUserRepository(db), validator}
}

func (h *handler) validate(data any) []response.FieldValidation {
	validations := []response.FieldValidation{}

	if errs := h.validator.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fieldJSONName := err.Field()

			field, ok := reflect.TypeOf(data).Elem().FieldByName(err.Field())
			if ok {
				found, ok := field.Tag.Lookup("json")
				if ok {
					fieldJSONName = found
				}
			}

			validations = append(validations, response.FieldValidation{
				Field:         fieldJSONName,
				ReceivedValue: err.Value(),
				Message:       "Validation failed for: " + err.Tag(),
			})
		}
	}

	return validations
}

func (h *handler) bodyParseAndValidate(c *fiber.Ctx, dataPtr any) (bool, error) {
	if err := c.BodyParser(dataPtr); err != nil {
		return false, c.Status(http.StatusBadRequest).JSON(response.RequestMalformed("request body is malformed"))
	}
	if validations := h.validate(dataPtr); len(validations) > 0 {
		return false, c.Status(http.StatusBadRequest).JSON(response.RequestMalformedWithValidations(validations))
	}

	return true, nil
}
