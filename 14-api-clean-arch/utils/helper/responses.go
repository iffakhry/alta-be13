package helper

func FailedResponse(msg string) map[string]any {
	return map[string]any{
		"status":  "failed",
		"message": msg,
	}
}

func SuccessResponse(msg string) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": msg,
	}
}

func SuccessWithDataResponse(msg string, data any) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}
