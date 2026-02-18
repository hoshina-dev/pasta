package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DataSourceName string
	Port           string
	CORSOrigins    string
	RedisURL       string
	RedisPassword  string
	RedisDB        int
}

func Load() *Config {
	_ = godotenv.Load()

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return &Config{
		DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
		Port:           getEnvOrDefault("PORT", "8080"),
		CORSOrigins:    getEnvOrDefault("CORS_ORIGINS", "*"),
		RedisURL:       getEnvOrDefault("REDIS_URL", "localhost:6379"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RedisDB:        redisDB,
	}
}

func getEnvOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
