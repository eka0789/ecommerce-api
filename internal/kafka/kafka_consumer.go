package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic string, group string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: group,
		}),
	}
}

func (c *KafkaConsumer) Start() {
	go func() {
		for {
			msg, err := c.reader.ReadMessage(context.Background())
			if err == nil {
				log.Println("[Kafka] Order processed:", string(msg.Value))
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
