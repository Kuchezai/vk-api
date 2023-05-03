package http

import (
	v1 "vk-api/internal/transport/http/v1"

	"github.com/gin-gonic/gin"
)

type logger interface {
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
}

func StartNewServer(router *gin.Engine, l logger, uh *v1.UserHandler, port string) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/user/:id", uh.GetUser)
	router.GET("/user/:id/friends", uh.GetUserFriends)

	l.Info("Server starting!")
	router.Run(port)

}
