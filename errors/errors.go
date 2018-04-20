package errors

import (
	"fmt"
)

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func New(code int, message string, field string) error {
	return &Error{
		Code:    code,
		Message: message,
		Field:   field,
	}
}
