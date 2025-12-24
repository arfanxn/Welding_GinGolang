package middleware

import (
	"net/http"

	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	permissionRepository "github.com/arfanxn/welding/internal/module/permission/domain/repository"
	permissionService "github.com/arfanxn/welding/internal/module/permission/usecase/service"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/httperror"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AuthorizeMiddleware interface {
	RequirePermissionNames(requiredPermNames ...permissionEnum.PermissionName) gin.HandlerFunc
}

type authorizeMiddleware struct {
	userRepository       userRepository.UserRepository
	permissionRepository permissionRepository.PermissionRepository

	permissionService permissionService.PermissionService
}

type NewAuthorizeMiddlewareParams struct {
	fx.In

	UserRepository       userRepository.UserRepository
	PermissionRepository permissionRepository.PermissionRepository

	PermissionService permissionService.PermissionService
}

func NewAuthorizeMiddleware(
	params NewAuthorizeMiddlewareParams,
) (AuthorizeMiddleware, error) {
	return &authorizeMiddleware{
		userRepository:       params.UserRepository,
		permissionRepository: params.PermissionRepository,

		permissionService: params.PermissionService,
	}, nil
}

func (m *authorizeMiddleware) RequirePermissionNames(
	requiredPermNames ...permissionEnum.PermissionName,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPermissions := c.MustGet(contextkey.UserPermissionsKey).([]*entity.Permission)

		hasPermissions, err := m.permissionService.CheckByNames(userPermissions, requiredPermNames...)
		if err != nil {
			panic(err)
		}

		if !hasPermissions {
			httperror.Panic(http.StatusForbidden, "User tidak memiliki hak akses", nil)
		}

		c.Next()
	}
}
