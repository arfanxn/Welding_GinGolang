package di

import (
	"github.com/arfanxn/welding/internal/module/material_test_machine/infrastructure/policy"
	repositoryImpl "github.com/arfanxn/welding/internal/module/material_test_machine/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/material_test_machine/presentation/http"
	"github.com/arfanxn/welding/internal/module/material_test_machine/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"material_test_machine",
	fx.Provide(
		repositoryImpl.NewGormMaterialTestMachineRepository,
		policy.NewMaterialTestMachinePolicy,
		usecase.NewMaterialTestMachineUsecase,
		http.NewMaterialTestMachineHandler,
	),
)
