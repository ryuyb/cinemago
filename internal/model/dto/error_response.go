package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse defines a unified error response structure
type ErrorResponse struct {
	// Code is the error code that identifies the specific error type
	Code int `json:"code"`
	// Message is a user-friendly error description
	Message string `json:"message"`
	// Details is an optional field for providing additional context information
	Details any `json:"details,omitempty"`
}

// NewErrorResponse creates a new instance of ErrorResponse
func NewErrorResponseWithData(code int, message string, details any) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// NewErrorResponse creates a new instance of ErrorResponse
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface, returning the error message
func (e *ErrorResponse) Error() string {
	return e.Message
}

// MarshalJSON customizes JSON serialization to ensure fields are output as expected
func (e *ErrorResponse) MarshalJSON() ([]byte, error) {
	type Alias ErrorResponse
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	})
}

func BadRequest(message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}
