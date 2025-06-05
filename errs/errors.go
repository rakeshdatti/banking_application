package errs

import (
	"net/http"
)

type AppError struct {
	Code    int `json:",omitempty"`
	Message string `json: "message"`
}

// Error implements error.
func (a *AppError) Error() string {
	panic("unimplemented")
}


func (e *AppError) AsMessage() *AppError{
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
