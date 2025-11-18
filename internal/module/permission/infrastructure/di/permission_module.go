package di

import (
	repositoryImpl "github.com/arfanxn/welding/internal/module/permission/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/permission/presentation/http"
	"github.com/arfanxn/welding/internal/module/permission/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"permission",
	fx.Provide(
		repositoryImpl.NewGormPermissionRepository,
		usecase.NewPermissionUsecase,
		http.NewPermissionHandler,
	),
)
