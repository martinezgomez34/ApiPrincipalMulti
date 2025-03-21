package consumer

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("Error conectando a RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error creando canal de RabbitMQ", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}

func (r *RabbitMQ) DeclareQueue(queueName string) (amqp.Queue, error) {
	return r.Channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (r *RabbitMQ) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(
		queueName,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return nil, err
	}
	
	return msgs, nil
}

func (r *RabbitMQ) PublishMessage(queueName string, body []byte) error {
	return r.Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (r *RabbitMQ) Close() {
	r.Channel.Close()
	r.Conn.Close()
}