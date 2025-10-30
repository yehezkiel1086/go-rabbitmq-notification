package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		Rabbitmq *Rabbitmq
	}

	App struct {
		Name string
		Env  string
	}

	Rabbitmq struct {
		User string
		Pass string
		Host string
		Port string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	Rabbitmq := &Rabbitmq{
		User: os.Getenv("RABBITMQ_USER"),
		Pass: os.Getenv("RABBITMQ_PASS"),
		Host: os.Getenv("RABBITMQ_HOST"),
		Port: os.Getenv("RABBITMQ_PORT"),
	}

	return &Container{
		App: App,
		Rabbitmq: Rabbitmq,
	}, nil
}
