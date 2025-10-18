package middleware

import (
	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
)

func HttpErrorRecoveryMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(*errorutil.HttpError); ok {
					c.AbortWithStatusJSON(err.Code, response.NewBodyWithErrors(err.Code, err.Error(), err.Errors))
					return
				}

				// If it's not an HttpError, re-panic to let Gin handle it
				panic(r)
			}
		}()

		c.Next()
	}
}
