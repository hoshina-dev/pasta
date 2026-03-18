package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataSourceName         string
	Port                   string
	CORSOrigins            string
	S3Bucket               string
	S3Region               string
	S3BaseURL              string
	RabbitMQURL            string
	RabbitMQExchange       string
	RabbitMQRoutingKey     string
	OptimizationWebhookURL string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		DataSourceName:         os.Getenv("DATA_SOURCE_NAME"),
		Port:                   getEnvOrDefault("PORT", "8080"),
		CORSOrigins:            getEnvOrDefault("CORS_ORIGINS", "*"),
		S3Bucket:               os.Getenv("S3_BUCKET"),
		S3Region:               getEnvOrDefault("S3_REGION", "us-east-1"),
		S3BaseURL:              os.Getenv("S3_BASE_URL"),
		RabbitMQURL:            getEnvOrDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		RabbitMQExchange:       getEnvOrDefault("RABBITMQ_EXCHANGE", "optimization"),
		RabbitMQRoutingKey:     getEnvOrDefault("RABBITMQ_ROUTING_KEY", "3d.optimize"),
		OptimizationWebhookURL: os.Getenv("OPTIMIZATION_WEBHOOK_URL"),
	}
}

func getEnvOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
