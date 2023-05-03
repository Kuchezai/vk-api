package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"vk-api/config"
	"vk-api/internal/repo/webapi"
	"vk-api/internal/transport/http"
	v1 "vk-api/internal/transport/http/v1"
	"vk-api/internal/usecase"
	"vk-api/pkg/logger"
)

func Run(cfg *config.Config) {

	wa := webapi.NewUserWebAPI(cfg.VK.KeyAPI)
	userUsecase := usecase.NewUserUsecase(wa)

	logger, err := logger.New("logs.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	userHandler := v1.NewUserHandler(userUsecase, logger)

	router := gin.New()
	http.StartNewServer(router, logger, userHandler, cfg.HTTP.Port)

}
