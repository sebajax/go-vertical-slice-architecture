package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (r *AppError) Error() string {
	return fmt.Sprintf(r.Message)
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func NotFound(message string) *AppError {
	return NewAppError(http.StatusNotFound, message)
}

func BadRequest(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message)
}

func Unauthorized(message string) *AppError {
	return NewAppError(http.StatusUnauthorized, message)
}

func Forbidden(message string) *AppError {
	return NewAppError(http.StatusForbidden, message)
}

func InternalServerError() *AppError {
	return NewAppError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR")
}
