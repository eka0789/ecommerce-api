package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/legitdev/ecommerce-api/config"
	"github.com/legitdev/ecommerce-api/internal/cache"
	"github.com/legitdev/ecommerce-api/internal/handler"
	"github.com/legitdev/ecommerce-api/internal/kafka"
	"github.com/legitdev/ecommerce-api/internal/queue"
	"github.com/legitdev/ecommerce-api/internal/repository"

	_ "github.com/legitdev/ecommerce-api/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	cfg := config.LoadConfig()
	mongo := config.ConnectMongo(cfg)
	redisClient := cache.NewRedisClient(cfg.RedisAddr)
	orderRepo := repository.NewOrderMongoRepo(mongo)
	orderLogRepo := repository.NewOrderLogRepo(mongo)
	rmq := queue.NewRabbitMQPublisher(cfg.RabbitMQ)
	kafkaProducer := kafka.NewKafkaProducer(cfg.KafkaBrokers, "order.processed")
	kafkaConsumer := kafka.NewKafkaConsumer(cfg.KafkaBrokers, "order.processed", "email-group")
	kafkaConsumer.Start()

	consumer := queue.NewOrderConsumer(cfg.RabbitMQ, "order_queue", redisClient, kafkaProducer)
	consumer.Start()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewOrderHandler(orderRepo, redisClient, orderLogRepo, rmq)
	h.RegisterRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}