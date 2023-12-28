package helper

type successResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func NewSuccessResponse(code int, message string, data interface{}) successResponse {
	return successResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type errorResponse struct {
	Code    int
	Message string
	Error   interface{}
}

func NewErrorResponse(code int, message string, err interface{}) errorResponse {
	return errorResponse{
		Code:    code,
		Message: "Error :" + message,
		Error:   err,
	}
}
