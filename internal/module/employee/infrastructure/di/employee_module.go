package di

import (
	repositoryImpl "github.com/arfanxn/welding/internal/module/employee/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"employee",
	fx.Provide(
		repositoryImpl.NewGormEmployeeRepository,
	),
)
