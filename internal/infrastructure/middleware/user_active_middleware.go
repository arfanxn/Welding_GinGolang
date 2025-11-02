package middleware

import (
	"net/http"

	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
)

var _ Middleware = (*userActiveMiddleware)(nil)

type UserActiveMiddleware interface {
	Middleware
}

type userActiveMiddleware struct {
}

func NewUserActiveMiddleware() (UserActiveMiddleware, error) {
	return &userActiveMiddleware{}, nil
}

// MiddlewareFunc returns a Gin middleware handler function that handles user verification for protected routes.
func (m *userActiveMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(contextkey.UserKey).(*entity.User)

		// Check if the user account is active
		if !user.IsActive() {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "User tidak aktif, silahkan hubungi admin", nil))
		}

		c.Next()
	}
}
