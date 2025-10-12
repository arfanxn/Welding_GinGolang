package cmd

import (
	"context"
	"os"

	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/arfanxn/welding/internal/infrastructure/di"
	"github.com/arfanxn/welding/internal/infrastructure/logger"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var migrateCommand = &cli.Command{
	Name:  "migrate",
	Usage: "Run the migration",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		app := fx.New(
			di.Module,
			fx.Invoke(_migrate),
		)

		app.Run()

		return nil
	},
}

type _migrateParams struct {
	fx.In

	Config *config.Config
	Logger *logger.Logger
	GormDB *gorm.DB
}

func _migrate(params _migrateParams) error {
	params.Logger.Info("==========migrating database==========")
	defer os.Exit(0)

	db, err := params.GormDB.DB()
	if err != nil {
		params.Logger.Error("failed to get database connection", zap.Error(err))
		return err
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		params.Logger.Error("failed to create postgres driver", zap.Error(err))
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/infrastructure/database/postgres/migration", "postgres", driver)
	if err != nil {
		params.Logger.Error("failed to create migrate instance", zap.Error(err))
		return err
	}

	migrationType := "up"
	if len(os.Args) == 3 {
		migrationType = os.Args[2]
	}

	if migrationType == "down" || migrationType == "fresh" {
		err = m.Down()
		if err != nil {
			params.Logger.Error("failed to run migration down", zap.Error(err))
			return err
		}

		params.Logger.Info("migration down completed")
	}

	if migrationType != "down" {
		err = m.Up()
		if err != nil {
			if err == migrate.ErrNoChange {
				params.Logger.Error("no migration up to run", zap.Error(err))
				return err
			}

			params.Logger.Error("failed to run migration up", zap.Error(err))
			return err
		}

		params.Logger.Info("migration up completed")
	}

	params.Logger.Info("==========database migrated==========")

	return nil
}
