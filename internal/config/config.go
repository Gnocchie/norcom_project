package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitmqHost string
	RabbitmqUser string
	RabbitmqPass string
	QueueName    string
	RootPath     string
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		RabbitmqHost: getenv("RABBITMQ_HOST", "localhost:5672"),
		RabbitmqUser: getenv("RABBITMQ_USER", "guest"),
		RabbitmqPass: getenv("RABBITMQ_PASS", "guest"),
		QueueName:    getenv("QUEUE_NAME", "file_events"),
		RootPath:     getenv("ROOT_PATH", "."),
	}

	log.Printf("Config loaded as: host=%s queue=%s root=%s\n",
		cfg.RabbitmqHost, cfg.QueueName, cfg.RootPath)

	return cfg
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
