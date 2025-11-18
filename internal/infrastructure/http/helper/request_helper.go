package helper

import (
	"net/http"
	"net/url"

	"github.com/arfanxn/welding/internal/infrastructure/http/request"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// MustBindValidate binds the request data to the provided request struct and validates it.
// It panics with an HTTP error if binding or validation fails.
//
// Parameters:
//   - c: Gin context containing the incoming HTTP request
//   - req: Request struct implementing the request.Request interface that will hold the bound data
//
// This function performs two main operations:
// 1. Binds the request data to the provided struct using Gin's ShouldBind
// 2. Validates the struct using the Validate() method from the request.Request interface
//
// If binding fails, it returns a 400 Bad Request error.
// If validation fails, it returns a 422 Unprocessable Entity error with detailed field errors.
func MustBindValidate(c *gin.Context, req request.Request) {
	// Step 1: Bind request data to the struct
	if err := c.ShouldBind(req); err != nil {
		err = httperror.New(http.StatusBadRequest, "Permintaan tidak valid. Silakan periksa kembali data yang dikirim.", nil)
		panic(err)
	}

	// Step 2: Validate the struct
	if err := req.Validate(); err != nil {
		// Check if the error is a validation error with multiple fields
		if validationErrs, ok := err.(validation.Errors); ok {
			var message string
			errsMap := make(map[string][]string)

			// Process each validation error
			for field, err := range validationErrs {
				errStr := err.Error()
				// Use the first error message as the main message
				if message == "" {
					message = errStr
				}
				// Store all errors in a map for detailed error reporting
				errsMap[field] = []string{errStr}
			}

			// Create and panic with a structured validation error
			err = httperror.New(http.StatusUnprocessableEntity, message, errsMap)
			panic(err)
		}

		// If it's not a validation error, panic with the original error
		panic(err)
	}
}

// URLFromC returns the URL of the current request
//
// Parameters:
//   - c: Gin context containing the incoming HTTP request
//
// Returns:
//   - url.URL: The URL of the current request
func URLFromC(c *gin.Context) url.URL {
	scheme := boolutil.Ternary(
		c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https",
		"https",
		"http")
	host := c.Request.Host
	path := c.Request.URL.Path
	rq := c.Request.URL.RawQuery

	u := url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     path,
		RawQuery: rq,
	}
	return u
}
