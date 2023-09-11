package main

import (
	"github.com/Arshia-Izadyar/Car-sale-api/src/api"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/cache"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := db.InitDB(cfg)
	if err != nil {
		logger.Fatal(err, logging.Postgres, logging.Startup, "cant postgres server", nil)
	}
	defer db.CloseDB()

	err = cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(err, logging.Redis, logging.Startup, "cant redis server", nil)

	}
	defer cache.CloseRedis()

	api.Init(cfg)

}
