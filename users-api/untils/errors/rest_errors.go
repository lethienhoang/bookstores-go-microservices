package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server",
	}
}
