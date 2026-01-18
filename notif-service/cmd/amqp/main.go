package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/handler"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/storage/rabbitmq"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/adapter/storage/rabbitmq/repository"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/core/service"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError(err, "failed to load .env configs")
	fmt.Println(".env configs loaded successfully")
	
	// init rabbitmq
	mq, err := rabbitmq.New(conf.RabbitMQ)
	handleError(err, "rabbitmq connection failed")
	fmt.Println("rabbitmq connected successfully")

	defer mq.Close()

	ctx := context.Background()

	// dependency injection
	notifRepo, err := repository.NewNotifRepository(mq)
	handleError(err, "unable to create queue")

	notifSvc := service.NewNotifService(notifRepo)
	notifHandler := handler.NewNotifHandler(notifSvc)

	// receive notification
	notifHandler.ReceiveNotif(ctx)
}
