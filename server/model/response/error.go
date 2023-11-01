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

func RequestMalformed(details string) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: "Request malformed",
			Details: details,
			ValidationErrors: nil,
		},
	}
}

func RequestMalformedWithValidations(validationErrors []FieldValidation) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: "Request malformed",
			Details: "validation failed",
			ValidationErrors: nil,
		},
	}
}

func InternalProblem(details string) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: "Server encountered a problem while processing the request",
			Details: details,
			ValidationErrors: nil,
		},
	}
}

