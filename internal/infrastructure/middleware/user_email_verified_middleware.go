package middleware

import (
	"net/http"

	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
)

var _ Middleware = (*userEmailVerifiedMiddleware)(nil)

type UserEmailVerifiedMiddleware interface {
	Middleware
}

type userEmailVerifiedMiddleware struct {
}

func NewUserEmailVerifiedMiddleware() (UserEmailVerifiedMiddleware, error) {
	return &userEmailVerifiedMiddleware{}, nil
}

// MiddlewareFunc returns a Gin middleware handler function that handles user verification for protected routes.
func (m *userEmailVerifiedMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(contextkey.UserKey).(*entity.User)

		// Check if the user email is verified
		if !user.IsEmailVerified() {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Email belum terverifikasi, silahkan verifikasi email Anda", nil))
		}

		c.Next()
	}
}
