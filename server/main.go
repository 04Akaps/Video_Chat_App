package main

import (
	"flag"
	"github.com/04Akaps/Video_Chat_App/pkg/app"
	"log"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cfg := app.NewConfig(*configFlag)

	app.NewApp(cfg)
}
