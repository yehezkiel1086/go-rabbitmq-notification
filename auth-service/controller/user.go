package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/model"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/storage/postgres"
)

type UserController struct {
	db *postgres.DB
	ch *amqp.Channel
	q  *amqp.Queue
}

func InitUserController(db *postgres.DB, ch *amqp.Channel, q *amqp.Queue) (*UserController) {
	return &UserController{
		db: db,
		ch: ch,
		q: q,
	}
}

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

func (u *UserController) Register(c *gin.Context) {
	var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create user in database
		user := model.User{Email: req.Email, Name: req.Name}
		if err := u.db.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		// Publish event to RabbitMQ
		body, err := json.Marshal(req)
		if err != nil {
			log.Printf("Failed to marshal message: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create message"})
			return
		}
		if err := u.ch.Publish(
			"",     // exchange
			u.q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		); err != nil {
			log.Printf("Failed to publish message: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
			return
		}

		log.Printf("✅ Published new user: %s", body)
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
