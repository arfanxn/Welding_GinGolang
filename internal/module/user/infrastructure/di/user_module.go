package di

import (
	// userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	userRepositoryImpl "github.com/arfanxn/welding/internal/module/user/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/user/presentation/http"
	"github.com/arfanxn/welding/internal/module/user/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		userRepositoryImpl.NewGormUserRepository,
		usecase.NewUserUsecase,
		http.NewUserHandler,
	),
)
