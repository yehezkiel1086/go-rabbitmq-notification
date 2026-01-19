package repository

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/storage/rabbitmq"
)

type NotifRepository struct {
	mq *rabbitmq.Rabbitmq
	q *amqp.Queue
}

func NewNotifRepository(mq *rabbitmq.Rabbitmq) (*NotifRepository, error) {
	// init queue
	q, err := mq.DeclareQueue("email_confirm")
	if err != nil {
		return nil, err
	}

	return &NotifRepository{
		mq: mq,
		q: q,
	}, nil
}

func (n *NotifRepository) ReceiveNotif(ctx context.Context) (<-chan amqp.Delivery, error) {
	return n.mq.Consume(n.q)
}
