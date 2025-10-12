package cmd

import (
	"context"
	"os"

	"github.com/arfanxn/welding/internal/infrastructure/database/postgres/seeder"
	"github.com/arfanxn/welding/internal/infrastructure/di"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var seedCommand = &cli.Command{
	Name:  "seed",
	Usage: "Run the seed",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		app := fx.New(
			di.Module,
			fx.Invoke(seed),
		)

		app.Run()

		return nil
	},
}

type seedParams struct {
	fx.In

	GormDB *gorm.DB
	Logger *logger.Logger
}

func seed(params seedParams) error {
	params.Logger.Info("==========seeding database==========")
	defer os.Exit(0)

	seeders := []seeder.Seeder{
		seeder.NewUserSeeder(params.GormDB),
	}

	for _, s := range seeders {
		seederName := reflectutil.GetStructName(s)
		params.Logger.Info("seeding " + seederName)

		err := s.Seed()
		if err != nil {
			params.Logger.Error(
				"failed to seed "+seederName,
				zap.Error(err),
				zap.String("seeder", seederName),
			)

			return err
		}
	}

	params.Logger.Info("==========database seeded==========")

	return nil
}
