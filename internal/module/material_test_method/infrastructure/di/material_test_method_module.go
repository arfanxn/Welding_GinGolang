package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_method/infrastructure/policy"
	repositoryImpl "github.com/arfanxn/welding/internal/module/material_test_method/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_method/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_method/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_method",
	fx.Provide(
		repositoryImpl.NewGormMaterialTestMethodRepository,
		policy.NewMaterialTestMethodPolicy,
		usecase.NewMaterialTestMethodUsecase,
		http.NewMaterialTestMethodHandler,
	),
)
