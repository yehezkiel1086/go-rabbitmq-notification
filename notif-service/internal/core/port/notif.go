package port

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type NotifRepository interface {
	ReceiveNotif(ctx context.Context) (<-chan amqp.Delivery, error)
}

type NotifService interface {
	ReceiveNotif(ctx context.Context) (<-chan amqp.Delivery, error)
}
