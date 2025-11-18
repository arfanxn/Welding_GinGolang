package di

import (
	"github.com/arfanxn/welding/internal/module/role_user/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("role_user",
	fx.Provide(
		repository.NewGormRoleUserRepository,
	),
)
