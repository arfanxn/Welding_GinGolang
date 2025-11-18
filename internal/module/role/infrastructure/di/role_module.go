package di

import (
	"github.com/arfanxn/welding/internal/module/role/infrastructure/policy"
	roleRepositoryImpl "github.com/arfanxn/welding/internal/module/role/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/role/presentation/http"
	"github.com/arfanxn/welding/internal/module/role/usecase"
	"github.com/arfanxn/welding/internal/module/role/usecase/step"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"role",
	fx.Provide(
		roleRepositoryImpl.NewGormRoleRepository,
		policy.NewRolePolicy,
		step.NewStoreRoleStep,
		step.NewUpdateRoleStep,
		usecase.NewRoleUsecase,
		http.NewRoleHandler,
	),
)
