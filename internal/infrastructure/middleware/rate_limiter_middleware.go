package middleware

import (
	"net/http"

	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type RateLimiterMiddleware interface {
	Middleware
}

type rateLimiterMiddleware struct {
	limiter *rate.Limiter
}

func NewRateLimiterMiddleware() RateLimiterMiddleware {
	// Allow 4 requests per second with a burst of 4
	return &rateLimiterMiddleware{
		limiter: rate.NewLimiter(rate.Limit(4), 4),
	}
}

func (r *rateLimiterMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !r.limiter.Allow() {
			httperror.Panic(http.StatusTooManyRequests, "Terlalu banyak permintaan", nil)
		}

		c.Next()
	}
}
