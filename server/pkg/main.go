package main

import (
	"flag"
	"github.com/04Akaps/Video_Chat_App/config"
	"github.com/04Akaps/Video_Chat_App/pkg/app"
	"log"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cfg := config.NewConfig(*configFlag)

	if app, err := app.NewApp(cfg); err != nil {
		panic(err)
	} else {
		go app.Wait()
		app.Run()
	}

}
