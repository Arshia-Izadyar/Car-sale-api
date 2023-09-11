package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var logger = logging.NewLogger(config.GetConfig())

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	err := redisClient.Close()
	if err != nil {
		logger.Fatal(err, logging.Redis, logging.Api, "cant close redis", nil)

	}

}

func Set[T any](key string, value T, duration time.Duration, c *redis.Client) error {
	v, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	// _, err = c.Set(context.Background(), key, v, duration).Result()
	// if err != nil {
	// 	return err
	// }
	c.Set(context.Background(), key, v, duration)
	return nil
}

func Get[T any](key string, c *redis.Client) (*T, error) {
	var dest = new(T)
	res, err := c.Get(context.Background(), key).Result()
	if err != nil {
		return dest, err
	}

	err = json.Unmarshal([]byte(res), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
