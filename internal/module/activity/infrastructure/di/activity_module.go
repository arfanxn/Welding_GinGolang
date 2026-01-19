package di

import (
	activityRepositoryImpl "github.com/arfanxn/welding/internal/module/activity/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/activity/presentation/http"
	"github.com/arfanxn/welding/internal/module/activity/presentation/http/presenter"
	"github.com/arfanxn/welding/internal/module/activity/usecase"
	"github.com/arfanxn/welding/internal/module/activity/usecase/service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"activity",
	fx.Provide(
		activityRepositoryImpl.NewGormActivityRepository,
		service.NewActivityService,
		usecase.NewActivityUsecase,
		presenter.NewActivityPresenter,
		http.NewActivityHandler,
	),
)
