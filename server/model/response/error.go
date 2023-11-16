package response

type FieldValidation struct {
	Field         string `json:"field"`
	ReceivedValue any    `json:"received_value"`
	Message       string `json:"message"`
}

func RequestMalformed(details string) map[string]any {
	return map[string]any{
		"message": "request malformed",
		"details": details,
	}
}

func RequestMalformedWithValidations(validations []FieldValidation) map[string]any {
	return map[string]any{
		"message":     "request malformed",
		"details":     "validation failed",
		"validations": validations,
	}
}

func InternalProblem(details string) map[string]any {
	return map[string]any{
		"message": "server encountered a problem while processing the request",
		"details": details,
	}
}
