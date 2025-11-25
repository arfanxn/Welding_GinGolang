package cmd

import (
	"context"
	"os"

	"github.com/arfanxn/welding/internal/infrastructure/database/factory"
	"github.com/arfanxn/welding/internal/infrastructure/database/seeder"
	"github.com/arfanxn/welding/internal/infrastructure/di"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/arfanxn/welding/pkg/typeutil"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var seedCommand = &cli.Command{
	Name:  "seed",
	Usage: "Run the seed",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		app := fx.New(
			di.Module,
			fx.Provide(
				fx.Annotate(factory.NewEmployeeFactory, fx.ResultTags(`name:"employee_factory"`)),
				fx.Annotate(factory.NewPermissionRoleFactory, fx.ResultTags(`name:"permission_role_factory"`)),
				fx.Annotate(factory.NewRoleFactory, fx.ResultTags(`name:"role_factory"`)),
				fx.Annotate(factory.NewRoleUserFactory, fx.ResultTags(`name:"role_user_factory"`)),
				fx.Annotate(factory.NewUserFactory, fx.ResultTags(`name:"user_factory"`)),
				fx.Annotate(factory.NewCodeFactory, fx.ResultTags(`name:"code_factory"`)),
				fx.Annotate(factory.NewActivityFactory, fx.ResultTags(`name:"activity_factory"`)),

				fx.Annotate(seeder.NewUserSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewRoleSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewPermissionSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewCodeSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewMaterialTestMethodSeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
				fx.Annotate(seeder.NewActivitySeeder, fx.As(new(seeder.Seeder)), fx.ResultTags(`group:"seeders"`)),
			),
			fx.WithLogger(func(l *logger.Logger) fxevent.Logger {
				fxLogger := l.With(zap.String("component", "fx"))
				return &fxevent.ZapLogger{Logger: fxLogger}
			}),
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
		mustGetByTypeFromSeeders[*seeder.CodeSeeder](seeders),
		mustGetByTypeFromSeeders[*seeder.MaterialTestMethodSeeder](seeders),
		mustGetByTypeFromSeeders[*seeder.ActivitySeeder](seeders),
	}
	return orderedSeeders
}

func seed(params seedParams) (err error) {
	// Handle any panic that might occur during seeding
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			params.Logger.Error("panic occurred during seeding",
				zap.Error(err),
				zap.Stack("stack"),
			)
		}
	}()

	params.Logger.Info("==========seeding database==========")
	defer os.Exit(0)

	orderedSeeders := mustGetOrderedSeeders(params.Seeders)

	for _, s := range orderedSeeders {
		seederName := reflectutil.GetStructName(s)
		params.Logger.Info("seeding " + seederName)

		// Nested defer with named return to capture errors from panics in Seed()
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = r.(error)
					params.Logger.Error("panic in seeder "+seederName,
						zap.Error(err),
						zap.String("seeder", seederName),
						zap.Stack("stack"),
					)
				}
			}()

			if err = s.Seed(); err != nil {
				params.Logger.Error(
					"failed to seed "+seederName,
					zap.Error(err),
					zap.String("seeder", seederName),
				)
			}
		}()

		// If there was an error in the seeder, return it
		if err != nil {
			return err
		}
	}

	params.Logger.Info("==========database seeded==========")

	return nil
}
