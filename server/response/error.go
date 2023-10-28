package response

import (
	"fmt"
)

type ErrorResponse struct {
	ErrorData errorResponseContent `json:"error"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprint(e.ErrorData)
}

func (e *ErrorResponse) AddValidation(validation ErrorResponseValidation) {
	e.ErrorData.ValidationErrors = append(e.ErrorData.ValidationErrors,
		validation)
}

type ErrorResponseValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type errorResponseContent struct {
	Message          string                    `json:"message"`
	Details          string                    `json:"details"`
	ValidationErrors []ErrorResponseValidation `json:"validationErrors"`
}

func SError(message string) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: message,
		},
	}
}

func SErrorFromErr(message string, err error) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: message,
			Details: err.Error(),
		},
	}
}

func Error(message string, details string, validationErrors []ErrorResponseValidation) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			message,
			details,
			validationErrors,
		},
	}
}

func NewValidation(field string, message string) ErrorResponseValidation {
	return ErrorResponseValidation{
		Field:   field,
		Message: message,
	}
}
