package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/middleware"
	userHttp "github.com/arfanxn/welding/internal/module/user/presentation/http"
	"github.com/arfanxn/welding/pkg/response"
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
	UserHandler userHttp.UserHandler
}

func RegisterRoutes(params RegisterRoutesParams) error {
	// API v1
	apiV1 := params.Router.Group("/api/v1")
	apiV1.Use(middleware.HttpErrorRecoveryMiddlewareFunc())
	apiV1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Body{
			Code:    http.StatusOK,
			Status:  response.StatusSuccess,
			Message: "OK",
		})
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
	}

	return nil
}
