package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/config"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/controller"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/messaging/rabbitmq"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/model"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/router"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/storage/postgres"
)

func main() {
	// get .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs loaded successfully")

	// init context
	ctx := context.Background()

	// init postgres config
	db, err := postgres.InitPostgres(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Postgres connected successfully")

	// migrate db
	if err := db.Migrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("✅ Database migrated successfully")

	// init rabbitmq
	mq, err := rabbitmq.InitRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	fmt.Println("✅ RabbitMQ connected successfully")

	ch, err := mq.DeclareChannel()
	if err != nil {
		log.Fatalf("Failed to declare channel: %v", err)
	}
	defer ch.Close()

	q, err := mq.DeclareQueue(ch, "user_registered")
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// init controllers
	userController := controller.InitUserController(db, ch, q)

	// init router
	r := router.InitRouter(
		userController,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Router initialized successfully")

	// serve api
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}
