package di

// TODO: delete this file
/*
import (
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	repositoryImpl "github.com/arfanxn/welding/internal/module/user/infrastructure/repository"

	userHttp "github.com/arfanxn/welding/internal/module/user/presentation/http"
	"github.com/arfanxn/welding/internal/module/user/usecase"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

// ProvideUserDeps provides user dependencies
func ProvideUserDeps(injector do.Injector) error {
	do.Provide(injector, func(i do.Injector) (repository.UserRepository, error) {
		gormDB := do.MustInvoke[*gorm.DB](i)
		return repositoryImpl.NewGormUserRepository(gormDB), nil
	})
	do.Provide(injector, func(i do.Injector) (usecase.UserUsecase, error) {
		userRepo := do.MustInvoke[repository.UserRepository](i)
		return usecase.NewUserUsecase(userRepo), nil
	})
	do.Provide(injector, func(i do.Injector) (userHttp.UserHandler, error) {
		userUsecase := do.MustInvoke[usecase.UserUsecase](i)
		return userHttp.NewUserHandler(userUsecase), nil
	})
	return nil
}

*/
