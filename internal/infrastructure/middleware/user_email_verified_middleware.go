package middleware

import (
	"net/http"

	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/gin-gonic/gin"
)

var _ Middleware = (*userEmailVerifiedMiddleware)(nil)

type UserEmailVerifiedMiddleware interface {
	Middleware
}

type userEmailVerifiedMiddleware struct {
	userRepository userRepository.UserRepository
}

func NewUserEmailVerifiedMiddleware(
	userRepository userRepository.UserRepository,
) (UserEmailVerifiedMiddleware, error) {
	return &userEmailVerifiedMiddleware{
		userRepository: userRepository,
	}, nil
}

// MiddlewareFunc returns a Gin middleware handler that enforces email verification for protected routes.
// This middleware implements role-based access control with the following rules:
// - SuperAdmin users bypass email verification requirements
// - All other users must have a verified email to proceed
// - Unauthorized requests will receive a 401 Unauthorized response
//
// The middleware expects the authenticated user to be available in the request context
// under the contextkey.UserKey. The user's roles are checked to determine if they have
// SuperAdmin privileges, which grants them an exception to the email verification requirement.
func (m *userEmailVerifiedMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the authenticated user from the request context
		user := c.MustGet(contextkey.UserKey).(*entity.User)

		// Check if user has SuperAdmin role
		isSuperAdmin, err := m.userRepository.HasRoleNames(user, []roleEnum.RoleName{roleEnum.SuperAdmin})
		if err != nil {
			panic(err) // Panic on repository errors as they indicate system issues
		}

		// Enforce email verification for non-SuperAdmin users
		if !isSuperAdmin && !user.IsEmailVerified() {
			httperror.Panic(
				http.StatusUnauthorized,
				"Email belum terverifikasi, silahkan verifikasi email Anda",
				nil,
			)
		}

		// Proceed to the next handler if all checks pass
		c.Next()
	}
}
