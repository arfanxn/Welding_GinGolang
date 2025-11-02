package middleware

import (
	"net/http"

	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AuthorizeMiddleware interface {
	RequirePermissionNames(requiredPermNames ...permissionEnum.PermissionName) gin.HandlerFunc
}

type authorizeMiddleware struct {
	userRepository userRepository.UserRepository
}

type NewAuthorizeMiddlewareParams struct {
	fx.In

	UserRepository userRepository.UserRepository
}

func NewAuthorizeMiddleware(
	params NewAuthorizeMiddlewareParams,
) (AuthorizeMiddleware, error) {
	return &authorizeMiddleware{
		userRepository: params.UserRepository,
	}, nil
}

func (m *authorizeMiddleware) RequirePermissionNames(
	requiredPermNames ...permissionEnum.PermissionName,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(contextkey.UserKey).(*entity.User)

		hasPermissions, err := m.userRepository.HasPermissionNames(user, requiredPermNames)
		if err != nil {
			panic(err)
		}

		if !hasPermissions {
			panic(errorutil.NewHttpError(http.StatusForbidden, "User tidak memiliki hak akses", nil))
		}

		c.Next()
	}
}
