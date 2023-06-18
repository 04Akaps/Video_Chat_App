package main

import (
	"flag"
	"github.com/04Akaps/Video_Chat_App/pkg/app"
	"log"
)

//
//github.com/gin-gonic/gin v1.9.1
//github.com/naoina/toml v0.1.1

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cfg := app.NewConfig(*configFlag)

	app.NewApp(cfg)
}
