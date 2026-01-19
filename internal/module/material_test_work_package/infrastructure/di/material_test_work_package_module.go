package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_work_package/infrastructure/policy"
	materialTestWorkPackageRepositoryImpl "github.com/arfanxn/welding/internal/module/material_test_work_package/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_work_package/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_work_package",
	fx.Provide(
		materialTestWorkPackageRepositoryImpl.NewGormMaterialTestWorkPackageRepository,
		policy.NewMaterialTestWorkPackagePolicy,
		usecase.NewMaterialTestWorkPackageUsecase,
		http.NewMaterialTestWorkPackageHandler,
	),
)
