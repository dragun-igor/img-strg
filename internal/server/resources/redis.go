package resources

import (
	"log"
	"time"

	"github.com/dragun-igor/img-strg/config"
	"github.com/dragun-igor/img-strg/pkg/redis"
)

// Инициализация редиса
func InitRedis(config *config.Config) (*redis.Client, error) {
	client, err := redis.NewClient(redis.Config{
		Host:              config.RedisHost,
		Port:              config.RedisPort,
		Password:          config.RedisPassword,
		ConnectionTimeout: time.Second * 5,
		OperationTimeout:  time.Second * 2,
	})
	if err != nil {
		return nil, err
	}
	log.Println("redis connected")
	return client, nil
}
