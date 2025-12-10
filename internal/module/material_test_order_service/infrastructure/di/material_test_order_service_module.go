package di

import (
	repositoryImpl "github.com/arfanxn/welding/internal/module/material_test_order_service/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("material_test_order_service",
	fx.Provide(
		repositoryImpl.NewGormMaterialTestOrderServiceRepository,
	),
)
