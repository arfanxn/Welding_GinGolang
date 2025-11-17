package di

import (
	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/arfanxn/welding/internal/infrastructure/database"
	"github.com/arfanxn/welding/internal/infrastructure/http"
	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/infrastructure/id"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/mail"
	"github.com/arfanxn/welding/internal/infrastructure/middleware"
	"github.com/arfanxn/welding/internal/infrastructure/security"
	codeDi "github.com/arfanxn/welding/internal/module/code/infrastructure/di"
	employeeDi "github.com/arfanxn/welding/internal/module/employee/infrastructure/di"
	permissionDi "github.com/arfanxn/welding/internal/module/permission/infrastructure/di"
	permissionRoleDi "github.com/arfanxn/welding/internal/module/permission_role/infrastructure/di"
	roleDi "github.com/arfanxn/welding/internal/module/role/infrastructure/di"
	roleUserDi "github.com/arfanxn/welding/internal/module/role_user/infrastructure/di"
	userDi "github.com/arfanxn/welding/internal/module/user/infrastructure/di"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Infrastructure
var Module = fx.Module("infrastructure",
	fx.Provide(
		// Core
		config.NewConfigFromEnv,
		database.NewPostgresGormDBFromConfig,
		logger.NewLoggerFromConfig,
		mail.NewSmtpMailServiceFromConfig,
		jwt.NewJWTServiceFromConfig,
		security.NewBcryptPasswordService,
		id.NewULIDIdService,
		http.NewRouterFromConfig,
		func(engine *gin.Engine) gin.IRouter { return engine },

		// Middleware(s)
		middleware.NewHttpErrorRecoveryMiddleware,
		middleware.NewRateLimiterMiddleware,
		middleware.NewAuthenticateMiddleware,
		middleware.NewAuthorizeMiddleware,
		middleware.NewUserActiveMiddleware,
		middleware.NewUserEmailVerifiedMiddleware,
	),

	// Modules
	userDi.Module,
	roleDi.Module,
	roleUserDi.Module,
	permissionDi.Module,
	permissionRoleDi.Module,
	employeeDi.Module,
	codeDi.Module,

	// Logger
	fx.WithLogger(func(logger *logger.Logger) fxevent.Logger {
		fxLogger := logger.With(zap.String("component", "fx"))
		return &fxevent.ZapLogger{Logger: fxLogger}
	}),

	// Invoke
	fx.Invoke(http.RegisterRoutes),
)
