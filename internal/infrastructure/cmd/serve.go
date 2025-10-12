package cmd

import (
	"context"
	"fmt"

	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/arfanxn/welding/internal/infrastructure/di"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var serveCommand = &cli.Command{
	Name:  "serve",
	Usage: "Run the application",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		app := fx.New(
			di.Module,
			fx.Invoke(serve),
		)

		app.Run()

		return nil
	},
}

type serveParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Logger    *logger.Logger
	Engine    *gin.Engine
	Config    *config.Config
}

func serve(params serveParams) {
	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				params.Logger.Info("starting application", zap.String("port", params.Config.AppPort))

				// Run in a goroutine and handle errors via channel
				go func() {
					if err := params.Engine.Run(fmt.Sprintf(":%s", params.Config.AppPort)); err != nil {
						params.Logger.Error("HTTP server error", zap.Error(err))
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				params.Logger.Info("application stopping")
				// Gin doesn't have a built-in shutdown, so we rely on process termination
				return nil
			},
		},
	)
}
