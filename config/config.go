package config

import (
	"os"
)

type Config struct {
	MongoURI     string
	RedisAddr    string
	RabbitMQ     string
	KafkaBrokers []string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI:     getEnv("MONGO_URI", "mongodb+srv://username:password@cluster0.kj0kwel.mongodb.net/"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		RabbitMQ:     getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		KafkaBrokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
