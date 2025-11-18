package di

import (
	"github.com/arfanxn/welding/internal/module/code/infrastructure/policy"
	codeRepositoryImpl "github.com/arfanxn/welding/internal/module/code/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/code/presentation/http"
	"github.com/arfanxn/welding/internal/module/code/usecase"
	"github.com/arfanxn/welding/internal/module/code/usecase/service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"code",
	fx.Provide(
		codeRepositoryImpl.NewGormCodeRepository,
		policy.NewCodePolicy,
		service.NewCodeService,
		usecase.NewCodeUsecase,
		http.NewCodeHandler,
	),
)
