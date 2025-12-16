package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_order_user/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("material_test_order_user",
	fx.Provide(
		repository.NewGormMaterialTestOrderUserRepository,
	),
)
