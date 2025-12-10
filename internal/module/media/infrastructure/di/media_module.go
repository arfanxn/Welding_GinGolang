package di

import (
	repositoryImpl "github.com/arfanxn/welding/internal/module/media/infrastructure/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("media",
	fx.Provide(repositoryImpl.NewGormMediaRepository),
)
