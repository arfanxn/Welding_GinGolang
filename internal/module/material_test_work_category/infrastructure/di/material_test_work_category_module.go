package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_work_category/infrastructure/policy"
	materialTestWorkCategoryRepositoryImpl "github.com/arfanxn/welding/internal/module/material_test_work_category/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_work_category/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_work_category",
	fx.Provide(
		materialTestWorkCategoryRepositoryImpl.NewGormMaterialTestWorkCategoryRepository,
		policy.NewMaterialTestWorkCategoryPolicy,
		usecase.NewMaterialTestWorkCategoryUsecase,
		http.NewMaterialTestWorkCategoryHandler,
	),
)
