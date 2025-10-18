package di

import (
	roleRepositoryImpl "github.com/arfanxn/welding/internal/module/role/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/role/presentation/http"
	"github.com/arfanxn/welding/internal/module/role/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"role",
	fx.Provide(
		roleRepositoryImpl.NewGormRoleRepository,
		usecase.NewRoleUsecase,
		http.NewRoleHandler,
	),
)
