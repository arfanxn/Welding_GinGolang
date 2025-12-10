package di

import (
	"github.com/arfanxn/welding/internal/module/customer/infrastructure/policy"
	customerRepositoryImpl "github.com/arfanxn/welding/internal/module/customer/infrastructure/repository"
	http "github.com/arfanxn/welding/internal/module/customer/presentation/http"
	usecase "github.com/arfanxn/welding/internal/module/customer/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"customer",
	fx.Provide(
		customerRepositoryImpl.NewGormCustomerRepository,
		policy.NewCustomerPolicy,
		usecase.NewCustomerUsecase,
		http.NewCustomerHandler,
	),
)
