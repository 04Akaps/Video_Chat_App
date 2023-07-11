package app

import (
	"github.com/04Akaps/Video_Chat_App/config"
	"github.com/04Akaps/Video_Chat_App/router"
	"log"
	"os"
)

type App struct {
	Config *config.Config
	router *router.Router
	stop   chan struct{}
}

func NewApp(cfg *config.Config) (*App, error) {

	a := &App{
		Config: cfg,
	}

	log.Println("Server Status : "+cfg.ServerInfo.Ip, cfg.ServerInfo.Port)

	a.router = router.NewRouter(cfg)

	return a, nil
}

func (a *App) Wait() {
	log.Println("Server started")
	<-a.stop
	os.Exit(1)
}

func (a *App) Run() {
	a.router.Run()
}
