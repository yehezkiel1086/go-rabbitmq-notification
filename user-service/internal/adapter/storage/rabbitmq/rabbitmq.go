package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/config"
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

func (mq *Rabbitmq) Send(q *amqp.Queue, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	return mq.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
  })
}

func (mq *Rabbitmq) SendJSON(q *amqp.Queue, data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	return mq.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:        data,
  })
}

func (mq *Rabbitmq) Close() {
	mq.conn.Close()
	mq.ch.Close()
}
