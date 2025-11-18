// Package response provides types and functions for generating standardized HTTP responses.
// It helps in maintaining consistent response structures across the application.
package response

import (
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/arfanxn/welding/pkg/numberutil"
)

// Common response status values
var (
	// StatusSuccess indicates a successful operation
	StatusSuccess = "success"
	// StatusError indicates an error occurred during the operation
	StatusError = "error"
)

// Body represents the standard response structure for HTTP responses
type Body struct {
	// Code is the HTTP status code
	Code int `json:"code"`
	// Status indicates whether the operation was successful or not
	Status string `json:"status"`
	// Message provides a human-readable message about the operation result
	Message string `json:"message"`
	// Errors contains validation or business logic error details, if any
	Errors map[string][]string `json:"errors,omitempty"`
	// Data contains the response payload, if any
	Data any `json:"data,omitempty"`
}

// getStatusFromCode determines the status string based on the HTTP status code.
// Returns StatusSuccess for 2xx status codes, StatusError otherwise.
func getStatusFromCode(code int) string {
	return boolutil.Ternary(numberutil.Between(code, 200, 299), StatusSuccess, StatusError)
}

// NewBody creates a new response body with the given status code and message.
// The status is automatically determined based on the status code.
func NewBody(code int, message string) *Body {
	return &Body{
		Code:    code,
		Status:  getStatusFromCode(code),
		Message: message,
	}
}

// NewBodyWithErrors creates a new error response body with the given status code,
// error message, and error details. The status is always set to StatusError.
func NewBodyWithErrors(code int, message string, errors map[string][]string) *Body {
	return &Body{
		Code:    code,
		Status:  StatusError,
		Message: message,
		Errors:  errors,
	}
}

// NewBodyWithData creates a new successful response body with the given status code,
// message, and data payload. The status is automatically determined based on the status code.
func NewBodyWithData(code int, message string, data any) *Body {
	return &Body{
		Code:    code,
		Status:  getStatusFromCode(code),
		Message: message,
		Data:    data,
	}
}
