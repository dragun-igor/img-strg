package resources

import (
	"log"
	"time"

	"github.com/dragun-igor/img-strg/config"
	"github.com/dragun-igor/img-strg/pkg/redis"
)

const (
	connectionTimeout = time.Second * 5
	operationTimeout  = time.Second * 2
)

// Инициализация редиса
func InitRedis(config *config.Config) (*redis.Client, error) {
	client, err := redis.NewClient(redis.Config{
		Host:              config.RedisHost,
		Port:              config.RedisPort,
		Password:          config.RedisPassword,
		ConnectionTimeout: connectionTimeout,
		OperationTimeout:  operationTimeout,
	})
	if err != nil {
		return nil, err
	}
	log.Println("redis connected")
	return client, nil
}
