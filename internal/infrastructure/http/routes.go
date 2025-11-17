package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/middleware"
	codeHttp "github.com/arfanxn/welding/internal/module/code/presentation/http"
	permissionEnum "github.com/arfanxn/welding/internal/module/permission/domain/enum"
	permissionHttp "github.com/arfanxn/welding/internal/module/permission/presentation/http"
	roleHttp "github.com/arfanxn/welding/internal/module/role/presentation/http"
	userHttp "github.com/arfanxn/welding/internal/module/user/presentation/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RegisterRoutesParams struct {
	fx.In

	// Router
	Router gin.IRouter

	// Utilities
	Logger *logger.Logger

	// Middlewares
	HttpErrorRecoveryMiddleware middleware.HttpErrorRecoveryMiddleware
	RateLimiterMiddleware       middleware.RateLimiterMiddleware
	AuthenticateMiddleware      middleware.AuthenticateMiddleware
	AuthorizeMiddleware         middleware.AuthorizeMiddleware
	UserActiveMiddleware        middleware.UserActiveMiddleware
	UserEmailVerifiedMiddleware middleware.UserEmailVerifiedMiddleware

	// Handlers
	UserHandler       userHttp.UserHandler
	RoleHandler       roleHttp.RoleHandler
	PermissionHandler permissionHttp.PermissionHandler
	CodeHandler       codeHttp.CodeHandler
}

func RegisterRoutes(params RegisterRoutesParams) error {
	// API v1
	apiV1 := params.Router.Group("/api/v1")
	apiV1.Use(
		params.HttpErrorRecoveryMiddleware.MiddlewareFunc(),
		params.RateLimiterMiddleware.MiddlewareFunc(),
	)
	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "OK"))
	})
	{
		// --------------------------------------------------
		// Public routes
		// --------------------------------------------------

		user := apiV1.Group("/users")
		user.POST("/", params.UserHandler.Login)
		user.POST("/login", params.UserHandler.Login)
		user.POST("/register", params.UserHandler.Register)
		user.POST("/verify-email", params.UserHandler.VerifyEmail)
		user.PATCH("/reset-password", params.UserHandler.ResetPassword)

		code := apiV1.Group("/codes")
		code.POST("/user-email-verification", params.CodeHandler.CreateUserEmailVerification)
		code.POST("/user-reset-password", params.CodeHandler.CreateUserResetPassword)
	}
	{
		// --------------------------------------------------
		// Protected routes
		// --------------------------------------------------

		requirePermissionName := params.AuthorizeMiddleware.RequirePermissionNames

		protected := apiV1.Group("")
		protected.Use(
			params.AuthenticateMiddleware.MiddlewareFunc(),
			params.UserActiveMiddleware.MiddlewareFunc(),
			params.UserEmailVerifiedMiddleware.MiddlewareFunc(),
		)

		user := protected.Group("/users")

		// Logout
		user.DELETE("/logout", params.UserHandler.Logout)

		// Me
		user.GET("/me", params.UserHandler.Me)
		user.PUT("/me", params.UserHandler.UpdateMeProfile)
		user.PATCH("/me/password", params.UserHandler.UpdateMePassword)

		// Users
		user.GET("", requirePermissionName(permissionEnum.UsersIndex), params.UserHandler.Paginate)
		user.GET("/:id", requirePermissionName(permissionEnum.UsersShow), params.UserHandler.Show)
		user.POST("", requirePermissionName(permissionEnum.UsersStore), params.UserHandler.Store)
		user.PUT("/:id", requirePermissionName(permissionEnum.UsersUpdate), params.UserHandler.Update)
		// ! Deprecated
		// user.PATCH("/:id/password", requirePermissionName(permissionEnum.UsersUpdate), params.UserHandler.UpdatePassword)
		user.PATCH("/:id/activation/toggle", requirePermissionName(permissionEnum.UsersUpdate), params.UserHandler.ToggleActivation)
		user.DELETE("/:id", requirePermissionName(permissionEnum.UsersDestroy), params.UserHandler.Destroy)

		// Roles
		role := protected.Group("/roles")
		role.GET("", requirePermissionName(permissionEnum.RolesIndex), params.RoleHandler.Paginate)
		role.GET("/:id", requirePermissionName(permissionEnum.RolesShow), params.RoleHandler.Find)
		role.POST("", requirePermissionName(permissionEnum.RolesStore), params.RoleHandler.Store)
		role.PUT("/:id", requirePermissionName(permissionEnum.RolesUpdate), params.RoleHandler.Update)
		role.PATCH("/:id/set-default", requirePermissionName(permissionEnum.RolesUpdate), params.RoleHandler.SetDefault)
		role.DELETE("/:id", requirePermissionName(permissionEnum.RolesDestroy), params.RoleHandler.Destroy)

		// Permissions
		permission := protected.Group("/permissions")
		permission.GET("", requirePermissionName(permissionEnum.PermissionsIndex), params.PermissionHandler.Paginate)

		// Codes
		code := protected.Group("/codes")
		code.POST("/user-register-invitation", requirePermissionName(permissionEnum.UsersStore), params.CodeHandler.CreateUserRegisterInvitation)

	}

	return nil
}
