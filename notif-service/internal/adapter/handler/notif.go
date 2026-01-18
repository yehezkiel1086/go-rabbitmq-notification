package handler

import (
	"context"
	"log"

	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/core/port"
)

type NotifHandler struct {
	svc port.NotifService
}


func NewNotifHandler(svc port.NotifService) *NotifHandler {
	return &NotifHandler{
		svc: svc,
	}
}

func (h *NotifHandler) ReceiveNotif(ctx context.Context) {
	msgs, err := h.svc.ReceiveNotif(ctx)
	if err != nil {
		return
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
