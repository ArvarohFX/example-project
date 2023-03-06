package db

import (
	"fmt"
	"github.com/ArvarohFX/example-project/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DBStorage struct {
	DB     *gorm.DB
	SiteDB *sqlx.DB
}

func NewDBStorage(cfg *config.Config, logger *zap.SugaredLogger, isSiteDBEnabled bool) (*DBStorage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Postgres.Addr,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
		cfg.Postgres.Port,
	)
	pgDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage for working with pgbouncer
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		return nil, err
	}
	return &DBStorage{
		DB: pgDB,
	}, nil
}
