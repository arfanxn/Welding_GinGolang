package database

import (
	"github.com/arfanxn/welding/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGormDBFromConfig(cfg *config.Config) (*gorm.DB, error) {
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), gormCfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}
