package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQPublisher(url string) *RabbitMQPublisher {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("order_queue", true, false, false, false, nil)
	return &RabbitMQPublisher{channel: ch, queue: q}
}

func (r *RabbitMQPublisher) PublishOrderID(orderID string) error {
	return r.channel.Publish("", r.queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(orderID),
	})
}