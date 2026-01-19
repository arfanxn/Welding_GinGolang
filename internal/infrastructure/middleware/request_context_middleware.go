package middleware

import (
	"context"

	"github.com/arfanxn/welding/internal/infrastructure/http/helper"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/gin-gonic/gin"
)

type RequestContextMiddleware interface {
	MiddlewareFunc() gin.HandlerFunc
}

type requestContextMiddleware struct {
}

func NewRequestContextMiddleware() RequestContextMiddleware {
	return &requestContextMiddleware{}
}

func (m *requestContextMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		for key, value := range map[contextkey.ContextKey]any{
			contextkey.ClientIpKey:   c.ClientIP(),
			contextkey.UserAgentKey:  c.Request.UserAgent(),
			contextkey.RequestURLKey: helper.URLFromC(c),
		} {
			c.Set(key, value)
			ctx = context.WithValue(ctx, key, value)
		}
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
