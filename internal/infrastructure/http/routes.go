package http

import (
	"net/http"

	"github.com/arfanxn/welding/internal/infrastructure/http/response"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/internal/infrastructure/middleware"
	activityHttp "github.com/arfanxn/welding/internal/module/activity/presentation/http"
	addressHttp "github.com/arfanxn/welding/internal/module/address/presentation/http"
	analyticHttp "github.com/arfanxn/welding/internal/module/analytic/presentation/http"
	codeHttp "github.com/arfanxn/welding/internal/module/code/presentation/http"
	customerHttp "github.com/arfanxn/welding/internal/module/customer/presentation/http"
	materialTestMachineHttp "github.com/arfanxn/welding/internal/module/material_test_machine/presentation/http"
	materialTestMethodHttp "github.com/arfanxn/welding/internal/module/material_test_method/presentation/http"
	materialTestOrderHttp "github.com/arfanxn/welding/internal/module/material_test_order/presentation/http"
	materialTestServiceHttp "github.com/arfanxn/welding/internal/module/material_test_service/presentation/http"
	materialTestWorkCategoryHttp "github.com/arfanxn/welding/internal/module/material_test_work_category/presentation/http"
	materialTestWorkPackageHttp "github.com/arfanxn/welding/internal/module/material_test_work_package/presentation/http"
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
	RequestContextMiddleware    middleware.RequestContextMiddleware
	RateLimiterMiddleware       middleware.RateLimiterMiddleware
	AuthenticateMiddleware      middleware.AuthenticateMiddleware
	AuthorizeMiddleware         middleware.AuthorizeMiddleware
	UserActiveMiddleware        middleware.UserActiveMiddleware
	UserEmailVerifiedMiddleware middleware.UserEmailVerifiedMiddleware

	// Handlers
	UserHandler                     userHttp.UserHandler
	RoleHandler                     roleHttp.RoleHandler
	PermissionHandler               permissionHttp.PermissionHandler
	CodeHandler                     codeHttp.CodeHandler
	ActivityHandler                 activityHttp.ActivityHandler
	MaterialTestMethodHandler       materialTestMethodHttp.MaterialTestMethodHandler
	MaterialTestMachineHandler      materialTestMachineHttp.MaterialTestMachineHandler
	MaterialTestServiceHandler      materialTestServiceHttp.MaterialTestServiceHandler
	MaterialTestWorkCategoryHandler materialTestWorkCategoryHttp.MaterialTestWorkCategoryHandler
	MaterialTestWorkPackageHandler  materialTestWorkPackageHttp.MaterialTestWorkPackageHandler
	MaterialTestOrderHandler        materialTestOrderHttp.MaterialTestOrderHandler
	AddressHandler                  addressHttp.AddressHandler
	CustomerHandler                 customerHttp.CustomerHandler
	AnalyticHandler                 analyticHttp.AnalyticHandler
}

func RegisterRoutes(params RegisterRoutesParams) error {

	router := params.Router
	router.Static("/storage", "./storage")

	// API v1
	apiV1 := router.Group("/api/v1")
	apiV1.Use(
		params.HttpErrorRecoveryMiddleware.MiddlewareFunc(),
		params.RequestContextMiddleware.MiddlewareFunc(),
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

		activity := apiV1.Group("/activities")
		activity.GET("/enums", params.ActivityHandler.EnumValues)
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
		role.GET("/:id", requirePermissionName(permissionEnum.RolesShow), params.RoleHandler.Show)
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

		// Activity
		activity := protected.Group("/activities")
		activity.GET("", requirePermissionName(permissionEnum.ActivitiesIndex), params.ActivityHandler.Index)
		activity.GET("/:id", requirePermissionName(permissionEnum.ActivitiesShow), params.ActivityHandler.Show)

		// Address
		address := protected.Group("/addresses")
		address.GET("", requirePermissionName(permissionEnum.AddressesIndex), params.AddressHandler.Paginate)
		address.GET("/:id", requirePermissionName(permissionEnum.AddressesShow), params.AddressHandler.Show)
		address.POST("", requirePermissionName(permissionEnum.AddressesStore), params.AddressHandler.Store)
		address.PUT("/:id", requirePermissionName(permissionEnum.AddressesUpdate), params.AddressHandler.Update)
		address.DELETE("/:id", requirePermissionName(permissionEnum.AddressesDestroy), params.AddressHandler.Destroy)

		// Customer
		customer := protected.Group("/customers")
		customer.GET("", requirePermissionName(permissionEnum.CustomersIndex), params.CustomerHandler.Paginate)
		customer.GET("/:id", requirePermissionName(permissionEnum.CustomersShow), params.CustomerHandler.Show)
		customer.POST("", requirePermissionName(permissionEnum.CustomersStore), params.CustomerHandler.Store)
		customer.PUT("/:id", requirePermissionName(permissionEnum.CustomersUpdate), params.CustomerHandler.Update)
		customer.DELETE("/:id", requirePermissionName(permissionEnum.CustomersDestroy), params.CustomerHandler.Destroy)

		// Material Test Method
		materialTestMethod := protected.Group("/material-test-methods")
		materialTestMethod.GET("", requirePermissionName(permissionEnum.MaterialTestMethodsIndex), params.MaterialTestMethodHandler.Paginate)
		materialTestMethod.GET("/:id", requirePermissionName(permissionEnum.MaterialTestMethodsShow), params.MaterialTestMethodHandler.Show)
		materialTestMethod.POST("", requirePermissionName(permissionEnum.MaterialTestMethodsStore), params.MaterialTestMethodHandler.Store)
		materialTestMethod.PUT("/:id", requirePermissionName(permissionEnum.MaterialTestMethodsUpdate), params.MaterialTestMethodHandler.Update)
		materialTestMethod.DELETE("/:id", requirePermissionName(permissionEnum.MaterialTestMethodsDestroy), params.MaterialTestMethodHandler.Destroy)

		// Material Test Machine
		materialTestMachine := protected.Group("/material-test-machines")
		materialTestMachine.GET("", requirePermissionName(permissionEnum.MaterialTestMachinesIndex), params.MaterialTestMachineHandler.Paginate)
		materialTestMachine.GET("/:id", requirePermissionName(permissionEnum.MaterialTestMachinesShow), params.MaterialTestMachineHandler.Show)
		materialTestMachine.POST("", requirePermissionName(permissionEnum.MaterialTestMachinesStore), params.MaterialTestMachineHandler.Store)
		materialTestMachine.PUT("/:id", requirePermissionName(permissionEnum.MaterialTestMachinesUpdate), params.MaterialTestMachineHandler.Update)
		materialTestMachine.DELETE("/:id", requirePermissionName(permissionEnum.MaterialTestMachinesDestroy), params.MaterialTestMachineHandler.Destroy)

		// Material Test Service
		materialTestService := protected.Group("/material-test-services")
		materialTestService.GET("", requirePermissionName(permissionEnum.MaterialTestServicesIndex), params.MaterialTestServiceHandler.Paginate)
		materialTestService.GET("/:id", requirePermissionName(permissionEnum.MaterialTestServicesShow), params.MaterialTestServiceHandler.Show)
		materialTestService.POST("", requirePermissionName(permissionEnum.MaterialTestServicesStore), params.MaterialTestServiceHandler.Store)
		materialTestService.PUT("/:id", requirePermissionName(permissionEnum.MaterialTestServicesUpdate), params.MaterialTestServiceHandler.Update)
		materialTestService.DELETE("/:id", requirePermissionName(permissionEnum.MaterialTestServicesDestroy), params.MaterialTestServiceHandler.Destroy)

		// Material Test Work Category
		materialTestWorkCategory := protected.Group("/material-test-work-categories")
		materialTestWorkCategory.GET("", requirePermissionName(permissionEnum.MaterialTestWorkCategoriesIndex), params.MaterialTestWorkCategoryHandler.Paginate)
		materialTestWorkCategory.GET("/:id", requirePermissionName(permissionEnum.MaterialTestWorkCategoriesShow), params.MaterialTestWorkCategoryHandler.Show)
		materialTestWorkCategory.POST("", requirePermissionName(permissionEnum.MaterialTestWorkCategoriesStore), params.MaterialTestWorkCategoryHandler.Store)
		materialTestWorkCategory.PUT("/:id", requirePermissionName(permissionEnum.MaterialTestWorkCategoriesUpdate), params.MaterialTestWorkCategoryHandler.Update)
		materialTestWorkCategory.DELETE("/:id", requirePermissionName(permissionEnum.MaterialTestWorkCategoriesDestroy), params.MaterialTestWorkCategoryHandler.Destroy)

		// Material Test Work Package
		materialTestWorkPackage := protected.Group("/material-test-work-packages")
		materialTestWorkPackage.GET("", requirePermissionName(permissionEnum.MaterialTestWorkPackagesIndex), params.MaterialTestWorkPackageHandler.Paginate)
		materialTestWorkPackage.GET("/:id", requirePermissionName(permissionEnum.MaterialTestWorkPackagesShow), params.MaterialTestWorkPackageHandler.Show)
		materialTestWorkPackage.POST("", requirePermissionName(permissionEnum.MaterialTestWorkPackagesStore), params.MaterialTestWorkPackageHandler.Store)
		materialTestWorkPackage.PUT("/:id", requirePermissionName(permissionEnum.MaterialTestWorkPackagesUpdate), params.MaterialTestWorkPackageHandler.Update)
		materialTestWorkPackage.DELETE("/:id", requirePermissionName(permissionEnum.MaterialTestWorkPackagesDestroy), params.MaterialTestWorkPackageHandler.Destroy)

		// Material Test Order
		materialTestOrder := protected.Group("/material-test-orders")
		materialTestOrder.GET("", params.MaterialTestOrderHandler.Paginate)
		materialTestOrder.GET("/:id", params.MaterialTestOrderHandler.Show)
		materialTestOrder.POST("", params.MaterialTestOrderHandler.Store)
		materialTestOrder.PUT("/:id", params.MaterialTestOrderHandler.Update)
		materialTestOrder.PATCH("/:id/approve", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.Approve)
		materialTestOrder.PATCH("/:id/reject", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.Reject)
		materialTestOrder.PATCH("/:id/cancel", params.MaterialTestOrderHandler.Cancel)
		materialTestOrder.PATCH("/:id/submit-payment", params.MaterialTestOrderHandler.SubmitPayment)
		materialTestOrder.PATCH("/:id/approve-payment", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.ApprovePayment)
		materialTestOrder.PATCH("/:id/reject-payment", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.RejectPayment)
		materialTestOrder.PATCH("/:id/test", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.Test)
		materialTestOrder.PATCH("/:id/refund", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.Refund)
		materialTestOrder.PATCH("/:id/complete", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.Complete)
		materialTestOrder.PUT("/:id/service-evaluations/:order_service_evaluation_id", requirePermissionName(permissionEnum.MaterialTestOrdersUpdate), params.MaterialTestOrderHandler.UpdateOrderServiceEvaluation)
		materialTestOrder.POST("/:id/medias", params.MaterialTestOrderHandler.StoreMedia)
		materialTestOrder.PUT("/:id/medias/:media_id", params.MaterialTestOrderHandler.UpdateMedia)
		materialTestOrder.DELETE("/:id/medias/:media_id", params.MaterialTestOrderHandler.DestroyMedia)

		// Analytics
		analytic := protected.Group("/analytics")
		analytic.GET("/summary", requirePermissionName(permissionEnum.AnalyticsIndex), params.AnalyticHandler.Summary)
		analytic.GET("/orders/trends", requirePermissionName(permissionEnum.AnalyticsIndex), params.AnalyticHandler.OrderTrends)
	}

	return nil
}
