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

type FieldValidation struct {
	Field         string `json:"field"`
	ReceivedValue any    `json:"received_value"`
	Message       string `json:"message"`
}

type errorResponseContent struct {
	Message          string            `json:"message"`
	Details          string            `json:"details"`
	ValidationErrors []FieldValidation `json:"validationErrors"`
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

func Error(message string, details string, validationErrors []FieldValidation) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			message,
			details,
			validationErrors,
		},
	}
}
