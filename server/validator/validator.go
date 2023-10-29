package validator

import (
	"github.com/go-playground/validator"
	"github.com/pewpewnor/portorico/server/response"
)

var validate = validator.New()

func Init() {

}

func Validate(data any) []response.FieldValidation {
	validations := []response.FieldValidation{}

	errs := validate.Struct(data)
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
