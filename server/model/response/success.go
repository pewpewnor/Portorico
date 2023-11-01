package response

type SuccessResponseData struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(message string, data any) SuccessResponseData {
	return SuccessResponseData{
		Message: message,
		Data:    data,
	}
}
