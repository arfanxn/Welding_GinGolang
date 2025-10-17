package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AuthorizeMiddleware struct {
}

type NewAuthorizeMiddlewareParams struct {
	fx.In
}

func NewAuthorizeMiddleware(
	params NewAuthorizeMiddlewareParams,
) (*AuthorizeMiddleware, error) {
	return &AuthorizeMiddleware{}, nil
}

// TODO: implement authorize middleware
func (m *AuthorizeMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
