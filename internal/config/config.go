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
	S3Endpoint             string
	RabbitMQURL            string
	RabbitMQExchange       string
	RabbitMQRoutingKey     string
	OptimizationWebhookURL string
}

func Load() *Config {
	_ = godotenv.Load()

	// Map S3_ACCESS_KEY to AWS_ACCESS_KEY_ID if not already set
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		if accessKey := os.Getenv("S3_ACCESS_KEY"); accessKey != "" {
			os.Setenv("AWS_ACCESS_KEY_ID", accessKey)
		}
	}

	// Map S3_SECRET_KEY to AWS_SECRET_ACCESS_KEY if not already set
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		if secretKey := os.Getenv("S3_SECRET_KEY"); secretKey != "" {
			os.Setenv("AWS_SECRET_ACCESS_KEY", secretKey)
		}
	}

	return &Config{
		DataSourceName:         os.Getenv("DATA_SOURCE_NAME"),
		Port:                   getEnvOrDefault("PORT", "8080"),
		CORSOrigins:            getEnvOrDefault("CORS_ORIGINS", "*"),
		S3Bucket:               os.Getenv("S3_BUCKET_NAME"),
		S3Region:               getEnvOrDefault("S3_REGION", "us-east-1"),
		S3BaseURL:              os.Getenv("S3_PUBLIC_URL"),
		S3Endpoint:             os.Getenv("S3_ENDPOINT"),
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
