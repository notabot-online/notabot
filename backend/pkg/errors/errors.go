package errors

import "fmt"

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	ErrNotFound      = &AppError{Code: "NOT_FOUND", Message: "Resource not found"}
	ErrInvalidInput  = &AppError{Code: "INVALID_INPUT", Message: "Invalid input"}
	ErrUnauthorized  = &AppError{Code: "UNAUTHORIZED", Message: "Unauthorized"}
	ErrInternalError = &AppError{Code: "INTERNAL_ERROR", Message: "Internal server error"}
)


