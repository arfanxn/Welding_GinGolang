package di

import (
	"github.com/arfanxn/welding/internal/module/permission_role/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("permission_role",
	fx.Provide(
		repository.NewGormPermissionRoleRepository,
	),
)
