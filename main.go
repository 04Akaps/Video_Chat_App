package main

import (
	"flag"
	"log"
	"streamingServer/pkg/app"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cfg := app.NewConfig(*configFlag)

	app.NewApp(cfg)
}
