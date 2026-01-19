package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_service/infrastructure/policy"
	repositoryImpl "github.com/arfanxn/welding/internal/module/material_test_service/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_service/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_service/usecase/step"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_service",
	fx.Provide(
		repositoryImpl.NewGormMaterialTestServiceRepository,
		step.NewStoreMaterialTestServiceStep,
		step.NewUpdateMaterialTestServiceStep,
		policy.NewMaterialTestServicePolicy,
		usecase.NewMaterialTestServiceUsecase,
		http.NewMaterialTestServiceHandler,
	),
)
