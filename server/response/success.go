package response

type SuccessResponseData struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SSuccess(message string) SuccessResponseData {
	return SuccessResponseData{
		Message: message,
	}
}

func Success(message string, data any) SuccessResponseData {
	return SuccessResponseData{
		Message: message,
		Data:    data,
	}
}
