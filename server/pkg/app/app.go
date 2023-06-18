package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Config *Config
	router *gin.Engine
}

func NewApp(cfg *Config) (*App, error) {
	router := gin.Default()
	a := &App{
		Config: cfg,
		router: router,
	}

	log.Println("Server Status : "+cfg.ServerInfo.Ip, cfg.ServerInfo.Port)

	router.Run(cfg.ServerInfo.Port)

	return a, nil
}
