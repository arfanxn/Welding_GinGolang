package di

import (
	"github.com/arfanxn/welding/internal/module/analytic/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/analytic/presentation/http"
	"github.com/arfanxn/welding/internal/module/analytic/presentation/http/presenter"
	"github.com/arfanxn/welding/internal/module/analytic/usecase"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repository.NewGormAnalyticRepository,
		usecase.NewAnalyticUsecase,
		presenter.NewSummaryPresenter,
		http.NewAnalyticHandler,
	),
)
