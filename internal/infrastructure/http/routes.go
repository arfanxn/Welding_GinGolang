package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/middleware"
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
	AuthenticateMiddleware *middleware.AuthenticateMiddleware

	// Handlers
	UserHandler       userHttp.UserHandler
	RoleHandler       roleHttp.RoleHandler
	PermissionHandler permissionHttp.PermissionHandler
}

func RegisterRoutes(params RegisterRoutesParams) error {
	// API v1
	apiV1 := params.Router.Group("/api/v1")
	apiV1.Use(middleware.HttpErrorRecoveryMiddlewareFunc())
	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.NewBody(http.StatusOK, "OK"))
	})
	{
		// --------------------------------------------------
		// Public routes
		// --------------------------------------------------

		apiV1.POST("/users/login", params.UserHandler.Login)
	}
	{
		// --------------------------------------------------
		// Protected routes
		// --------------------------------------------------

		protected := apiV1.Group("")
		protected.Use(params.AuthenticateMiddleware.MiddlewareFunc())

		user := protected.Group("/users")
		user.GET("/me", params.UserHandler.Me)
		user.DELETE("/logout", params.UserHandler.Logout)
		user.GET("", params.UserHandler.Paginate)
		user.GET("/:id", params.UserHandler.Find)
		user.POST("", params.UserHandler.Store)
		user.PUT("/:id", params.UserHandler.Update)
		user.DELETE("/:id", params.UserHandler.Destroy)

		role := protected.Group("/roles")
		role.GET("", params.RoleHandler.Paginate)
		role.GET("/:id", params.RoleHandler.Find)
		role.POST("", params.RoleHandler.Store)
		role.PUT("/:id", params.RoleHandler.Update)
		role.DELETE("/:id", params.RoleHandler.Destroy)

		permission := protected.Group("/permissions")
		permission.GET("", params.PermissionHandler.Paginate)

	}

	return nil
}
