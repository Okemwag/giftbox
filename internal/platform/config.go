package platform

import (
	"os"
	"time"
)

type Config struct {
	Environment        string
	HTTPAddr           string
	WebhookAddr        string
	DatabaseDSN        string
	WorkerPollInterval time.Duration
}

func LoadConfig() Config {
	return Config{
		Environment:        env("APP_ENV", "local"),
		HTTPAddr:           env("HTTP_ADDR", ":8080"),
		WebhookAddr:        env("WEBHOOK_ADDR", ":8081"),
		DatabaseDSN:        os.Getenv("DATABASE_DSN"),
		WorkerPollInterval: durationEnv("WORKER_POLL_INTERVAL", 5*time.Second),
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func durationEnv(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}
