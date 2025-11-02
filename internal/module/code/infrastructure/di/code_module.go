package di

import (
	"github.com/arfanxn/welding/internal/module/code/infrastructure/policy"
	codeRepositoryImpl "github.com/arfanxn/welding/internal/module/code/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/code/presentation/http"
	"github.com/arfanxn/welding/internal/module/code/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"code",
	fx.Provide(
		codeRepositoryImpl.NewGormCodeRepository,
		policy.NewCodePolicy,
		usecase.NewCodeUsecase,
		http.NewCodeHandler,
	),
)
