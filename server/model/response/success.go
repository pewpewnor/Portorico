package response

func Success(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}
