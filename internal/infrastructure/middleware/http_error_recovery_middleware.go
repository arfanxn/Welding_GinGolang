package middleware

import (
	"github.com/arfanxn/welding/pkg/errors"
	"github.com/arfanxn/welding/pkg/response"
	"github.com/gin-gonic/gin"
)

func HttpErrorRecoveryMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(*errors.HttpError); ok {
					body := response.Body{
						Code:    err.Code,
						Status:  response.StatusError,
						Message: err.Error(),
					}

					if err.Errors != nil {
						body.Errors = err.Errors
					}

					c.AbortWithStatusJSON(err.Code, body)
					return
				}

				// If it's not an HttpError, re-panic to let Gin handle it
				panic(r)
			}
		}()

		c.Next()
	}
}
