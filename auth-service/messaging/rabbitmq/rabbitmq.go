package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func InitRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return &RabbitMQ{}, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	return &RabbitMQ{
		conn: conn,
	}, nil
}

func (r *RabbitMQ) DeclareChannel() (*amqp.Channel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return &amqp.Channel{}, fmt.Errorf("failed to open channel: %v", err)
	}

	defer ch.Close()

	return ch, nil
}

func (r *RabbitMQ) DeclareQueue(ch *amqp.Channel, name string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name, // queue name
		true,              // durable
		false,             // auto delete
		false,             // exclusive
		false,             // no wait
		nil,               // arguments
	)
	if err != nil {
		return &amqp.Queue{}, fmt.Errorf("failed to declare queue: %v", err)
	}

	return &q, nil
}

func (r *RabbitMQ) GetConnection() *amqp.Connection {
	return r.conn
}