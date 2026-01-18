package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/config"
)

type Rabbitmq struct {
	conn *amqp.Connection
	ch *amqp.Channel
}

func New(conf *config.RabbitMQ) (*Rabbitmq, error) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port)

	// connect to rabbitmq server
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	// create channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Rabbitmq{
		conn,
		ch,
	}, nil
}

func (mq *Rabbitmq) DeclareQueue(name string) (*amqp.Queue, error) {
	q, err := mq.ch.QueueDeclare(
		name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (mq *Rabbitmq) Consume(q *amqp.Queue) (<-chan amqp.Delivery, error) {
	return mq.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}

func (mq *Rabbitmq) Close() {
	mq.conn.Close()
	mq.ch.Close()
}
