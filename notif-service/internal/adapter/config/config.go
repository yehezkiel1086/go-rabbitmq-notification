package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App      *App
		AMQP *AMQP
		RabbitMQ *RabbitMQ
	}

	App struct {
		Name string
		Env  string
	}

	AMQP struct {
		Host string
		Port string
	}

	RabbitMQ struct {
		Host     string
		Port     string
		User string
		Password string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	AMQP := &AMQP{
		Host: os.Getenv("AMQP_HOST"),
		Port: os.Getenv("AMQP_PORT"),
	}

	Rabbitmq := &RabbitMQ{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		User: os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}

	return &Container{
		App:      App,
		AMQP: AMQP,
		RabbitMQ: Rabbitmq,
	}, nil
}
