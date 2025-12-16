package di

import (
	repositoryImpl "github.com/arfanxn/welding/internal/module/media/infrastructure/repository"
	"github.com/arfanxn/welding/internal/module/media/usecase/service"
	"go.uber.org/fx"
)

var Module = fx.Module("media",
	fx.Provide(
		repositoryImpl.NewGormMediaRepository,
		service.NewMediaService,
	),
)
