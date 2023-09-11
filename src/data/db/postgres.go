package db

import (
	"fmt"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDB(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
	)
	var err error
	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := dbClient.DB()

	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)
	logger.Info(logging.Postgres, logging.Startup, "postgres started", nil)
	return nil
}

func GetDB() *gorm.DB { return dbClient }

func CloseDB() {
	db, err := dbClient.DB()
	if err != nil {
		logger.Fatal(err, logging.Postgres, logging.Api, "cant get postgres db", nil)
	}
	err = db.Close()
	if err != nil {
		logger.Fatal(err, logging.Postgres, logging.Api, "cant close postgres", nil)
	}
}
