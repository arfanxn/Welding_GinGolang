package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/policy"
	repositoryImpl "github.com/arfanxn/welding/internal/module/material_test_order/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_order/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase"
	"github.com/arfanxn/welding/internal/module/material_test_order/usecase/step"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_order",
	fx.Provide(
		repositoryImpl.NewGormMaterialTestOrderRepository,
		policy.NewMaterialTestOrderPolicy,
		step.NewStoreMaterialTestOrderStep,
		step.NewUpdateMaterialTestOrderStep,
		usecase.NewMaterialTestOrderUsecase,
		http.NewMaterialTestOrderHandler,
	),
)
