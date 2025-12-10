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
	activityDi "github.com/arfanxn/welding/internal/module/activity/infrastructure/di"
	addressDi "github.com/arfanxn/welding/internal/module/address/infrastructure/di"
	codeDi "github.com/arfanxn/welding/internal/module/code/infrastructure/di"
	customerDi "github.com/arfanxn/welding/internal/module/customer/infrastructure/di"
	employeeDi "github.com/arfanxn/welding/internal/module/employee/infrastructure/di"
	materialTestMachineDi "github.com/arfanxn/welding/internal/module/material_test_machine/infrastructure/di"
	materialTestMethodDi "github.com/arfanxn/welding/internal/module/material_test_method/infrastructure/di"
	materialTestOrderDi "github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/di"
	materialTestOrderServiceDi "github.com/arfanxn/welding/internal/module/material_test_order_service/infrastructure/di"
	materialTestOrderServiceEvaluationDi "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/infrastructure/di"
	materialTestServiceDi "github.com/arfanxn/welding/internal/module/material_test_service/infrastructure/di"
	materialTestWorkCategoryDi "github.com/arfanxn/welding/internal/module/material_test_work_category/infrastructure/di"
	materialTestWorkPackageDi "github.com/arfanxn/welding/internal/module/material_test_work_package/infrastructure/di"
	mediaDi "github.com/arfanxn/welding/internal/module/media/infrastructure/di"
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
		database.NewMongoClientFromConfig,
		database.NewMongoDatabaseFromClientAndConfig,
		logger.NewLoggerFromConfig,
		mail.NewSmtpMailServiceFromConfig,
		jwt.NewJWTServiceFromConfig,
		security.NewBcryptPasswordService,
		id.NewULIDIdService,
		http.NewRouterFromConfig,
		func(engine *gin.Engine) gin.IRouter { return engine },

		// Middleware(s)
		middleware.NewHttpErrorRecoveryMiddleware,
		middleware.NewRequestContextMiddleware,
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
	activityDi.Module,
	materialTestMethodDi.Module,
	materialTestMachineDi.Module,
	materialTestServiceDi.Module,
	materialTestWorkCategoryDi.Module,
	materialTestWorkPackageDi.Module,
	addressDi.Module,
	customerDi.Module,
	materialTestOrderDi.Module,
	materialTestOrderServiceDi.Module,
	materialTestOrderServiceEvaluationDi.Module,
	mediaDi.Module,

	// Logger
	fx.WithLogger(func(logger *logger.Logger) fxevent.Logger {
		fxLogger := logger.With(zap.String("component", "fx"))
		return &fxevent.ZapLogger{Logger: fxLogger}
	}),

	// Invoke
	fx.Invoke(http.RegisterRoutes),
)
