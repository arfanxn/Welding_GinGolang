package di

import (
	// userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/internal/module/user/infrastructure/policy"
	userRepositoryImpl "github.com/arfanxn/welding/internal/module/user/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/user/presentation/http"
	"github.com/arfanxn/welding/internal/module/user/usecase"
	"github.com/arfanxn/welding/internal/module/user/usecase/step"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		userRepositoryImpl.NewGormUserRepository,
		policy.NewUserPolicy,
		step.NewRegisterUserStep,
		step.NewSaveUserStep,
		usecase.NewUserUsecase,
		http.NewUserHandler,
	),
)
