package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ecommerce-api/config"
	"ecommerce-api/internal/cache"
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/kafka"
	"ecommerce-api/internal/queue"
	"ecommerce-api/internal/repository"

	_ "ecommerce-api/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	log.Println("🚀 Starting ecommerce-api...")

	// Load config from .env or env vars
	cfg := config.LoadConfig()

	// Connect to MongoDB
	var mongoClient = config.ConnectMongo(cfg)
	if mongoClient == nil {
		log.Fatal("❌ Failed to connect to MongoDB")
	}
	log.Println("✅ Connected to MongoDB")

	// Connect to Redis
	redisClient := cache.NewRedisClient(cfg.RedisAddr)
	if redisClient == nil {
		log.Fatal("❌ Failed to connect to Redis")
	}
	log.Println("✅ Connected to Redis")

	// Setup repositories
	orderRepo := repository.NewOrderMongoRepo(mongoClient)
	orderLogRepo := repository.NewOrderLogRepo(mongoClient)

	// Setup RabbitMQ publisher
	rmq := queue.NewRabbitMQPublisher(cfg.RabbitMQ)
	if rmq == nil {
		log.Fatal("❌ Failed to connect to RabbitMQ")
	}
	log.Println("✅ RabbitMQ publisher ready")

	// Setup Kafka producer
	kafkaProducer := kafka.NewKafkaProducer(cfg.KafkaBrokers, "order.processed")
	if kafkaProducer == nil {
		log.Fatal("❌ Failed to connect to Kafka producer")
	}
	log.Println("✅ Kafka producer ready")

	// Start Kafka consumer for logging/email simulation
	go kafka.NewKafkaConsumer(cfg.KafkaBrokers, "order.processed", "email-group").Start()

	// Start RabbitMQ order consumer
	go queue.NewOrderConsumer(cfg.RabbitMQ, "order_queue", redisClient, kafkaProducer).Start()

	// Setup Gin router
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup order routes
	h := handler.NewOrderHandler(orderRepo, redisClient, orderLogRepo, rmq)
	h.RegisterRoutes(r)

	// Start server
	addr := ":8080"
	log.Printf("🌐 Listening on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
