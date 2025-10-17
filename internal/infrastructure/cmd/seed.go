package cmd

import (
	"context"
	"os"

	"github.com/arfanxn/welding/internal/infrastructure/database/seeder"
	"github.com/arfanxn/welding/internal/infrastructure/di"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var seedCommand = &cli.Command{
	Name:  "seed",
	Usage: "Run the seed",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		app := fx.New(
			di.Module,
			fx.Provide(
				fx.Annotate(seeder.NewUserSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewRoleSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewPermissionSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
			),
			fx.Invoke(seed),
		)

		app.Run()

		return nil
	},
}

type seedParams struct {
	fx.In

	Logger  *logger.Logger
	Seeders []seeder.Seeder `group:"seeders"`
}

func mustGetByTypeFromSeeders[T any](seeders []seeder.Seeder) T {
	return typeutil.MustGetByTypeFromArray[seeder.Seeder, T](seeders)
}

func mustGetOrderedSeeders(seeders []seeder.Seeder) []seeder.Seeder {
	orderedSeeders := []seeder.Seeder{
		mustGetByTypeFromSeeders[*seeder.PermissionSeeder](seeders),
		mustGetByTypeFromSeeders[*seeder.RoleSeeder](seeders),
		mustGetByTypeFromSeeders[*seeder.UserSeeder](seeders),
	}
	return orderedSeeders
}

func seed(params seedParams) error {
	params.Logger.Info("==========seeding database==========")
	defer os.Exit(0)

	orderedSeeders := mustGetOrderedSeeders(params.Seeders)

	for _, s := range orderedSeeders {
		seederName := reflectutil.GetStructName(s)
		params.Logger.Info("seeding " + seederName)

		if err := s.Seed(); err != nil {
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
