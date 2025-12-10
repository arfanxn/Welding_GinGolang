package di

import (
	"github.com/arfanxn/welding/internal/module/address/infrastructure/policy"
	addressRepositoryImpl "github.com/arfanxn/welding/internal/module/address/infrastructure/repository"
	http "github.com/arfanxn/welding/internal/module/address/presentation/http"
	usecase "github.com/arfanxn/welding/internal/module/address/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"address",
	fx.Provide(
		addressRepositoryImpl.NewGormAddressRepository,
		policy.NewAddressPolicy,
		usecase.NewAddressUsecase,
		http.NewAddressHandler,
	),
)
