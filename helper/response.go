package helper

type SuccessResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func NewSuccessResponse(code int, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type ErrorResponse struct {
	Code    int
	Message string
	Error   interface{}
}

func NewErrorResponse(code int, message string, err interface{}) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: "Error : " + message,
		Error:   err,
	}
}
