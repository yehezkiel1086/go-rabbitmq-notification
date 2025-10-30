package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/email-service/config"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/email-service/messaging/rabbitmq"
)

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func main() {
	// get .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs loaded successfully")
	
	// init rabbitmq
	mq, err := rabbitmq.InitRabbitMQ(conf.Rabbitmq)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	conn := mq.GetConnection()
	defer conn.Close()

	// declare channel
	ch, err := mq.DeclareChannel()
	if err != nil {
		log.Fatalf("Failed to declare channel: %v", err)
	}
	defer ch.Close()

	// declare queue
	q, err := mq.DeclareQueue(ch, "user_registered")
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// consume messages
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("📨 Waiting for messages...")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var user UserResponse
			err := json.Unmarshal(d.Body, &user)
			if err != nil {
				log.Printf("Invalid message format: %v", err)
				continue
			}
			log.Printf("📧 Sending welcome email to %s <%s>", user.Name, user.Email)
		}
	}()

	<-forever
}
