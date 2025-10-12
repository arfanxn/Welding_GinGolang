package di

import (
	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/arfanxn/welding/internal/infrastructure/database/postgres"
	"github.com/arfanxn/welding/internal/infrastructure/http"
	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/middleware"
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
		postgres.NewPostgresGormDBFromConfig,
		logger.NewLoggerFromConfig,
		jwt.NewJWTServiceFromConfig,
		http.NewRouterFromConfig,
		func(engine *gin.Engine) gin.IRouter { return engine },

		// Middleware(s)
		middleware.NewAuthenticateMiddleware,
	),

	// Modules
	userDi.Module,

	// Logger
	fx.WithLogger(func(logger *logger.Logger) fxevent.Logger {
		fxLogger := logger.With(zap.String("component", "fx"))

		return &fxevent.ZapLogger{Logger: fxLogger}
	}),

	// Invoke
	fx.Invoke(http.RegisterRoutes),
)
