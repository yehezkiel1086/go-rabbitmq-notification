package service

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/core/port"
)

type NotifService struct {
	repo port.NotifRepository
}

func NewNotifService(repo port.NotifRepository) *NotifService {
	return &NotifService{
		repo: repo,
	}
}

func (s *NotifService) ReceiveNotif(ctx context.Context) (<-chan amqp.Delivery, error) {
	return s.repo.ReceiveNotif(ctx)
}
