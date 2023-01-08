package config

import "os"

const (
	defaultGRPCAddr    string = "localhost:50051"
	defaultRedisHost   string = "localhost"
	defaultRedisPort   string = "5432"
	defaultStoragePath string = "images"
)

type Config struct {
	GRPCAddr      string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	StoragePath   string
}

func getEnv(key, defaultValue string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return env
}

func Get() *Config {
	return &Config{
		GRPCAddr:      getEnv("GRPC_ADDR", defaultGRPCAddr),
		RedisHost:     getEnv("REDIS_HOST", defaultRedisHost),
		RedisPort:     getEnv("REDIS_PORT", defaultRedisPort),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		StoragePath:   getEnv("STORAGE_PATH", defaultStoragePath),
	}
}
