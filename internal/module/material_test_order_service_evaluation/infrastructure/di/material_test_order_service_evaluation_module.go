package di

import (
	materialTestOrderServiceEvaluationRepositoryImpl "github.com/arfanxn/welding/internal/module/material_test_order_service_evaluation/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_order_service_evaluation",
	fx.Provide(
		materialTestOrderServiceEvaluationRepositoryImpl.NewGormMaterialTestOrderServiceEvaluationRepository,
	),
)
