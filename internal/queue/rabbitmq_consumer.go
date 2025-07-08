package queue

import (
	"context"
	"log"
	"time"

	"github.com/legitdev/ecommerce-api/internal/cache"
	"github.com/legitdev/ecommerce-api/internal/kafka"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderConsumer struct {
	conn      *amqp.Connection
	queueName string
	cache     *cache.RedisClient
	kafka     *kafka.KafkaProducer
}

func NewOrderConsumer(url, queue string, redis *cache.RedisClient, kafka *kafka.KafkaProducer) *OrderConsumer {
	conn, _ := amqp.Dial(url)
	return &OrderConsumer{conn: conn, queueName: queue, cache: redis, kafka: kafka}
}

func (c *OrderConsumer) Start() {
	ch, _ := c.conn.Channel()
	msgs, _ := ch.Consume(c.queueName, "", true, false, false, false, nil)
	go func() {
		for m := range msgs {
			orderID := string(m.Body)
			_ = c.cache.SetStatus(context.Background(), orderID, "processed")
			_ = c.kafka.Publish("order.processed", []byte(orderID))
			log.Println("Order processed:", orderID)
			time.Sleep(1 * time.Second)
		}
	}()
}
