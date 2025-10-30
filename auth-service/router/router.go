package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/config"
	"github.com/yehezkiel1086/go-gin-rabbitmq-email-notif/auth-service/controller"
)

type Router struct {
	r *gin.Engine
}

func InitRouter(
	userController *controller.UserController,
) (*Router) {
	r := gin.Default()

	pb := r.Group("/api/v1")

	pb.POST("/register", userController.Register)

	return &Router{
		r: r,
	}
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}
