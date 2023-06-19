package router

import (
	"github.com/04Akaps/Video_Chat_App/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Router struct {
	engine *gin.Engine
	port   string
	rooms  *RoomMap
}

func NewRouter(cfg *config.Config) *Router {
	r := Router{
		engine: gin.New(),
		port:   cfg.ServerInfo.Port,
	}

	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())

	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.rooms = NewAllRooms()

	r.engine.POST("/create", r.CreateRoom)
	r.engine.POST("/join", r.JoinRoom)

	return &r
}

func (r *Router) Run() error {
	log.Println("Http Server Start", "endpoint", r.port)
	return r.engine.Run(r.port)
}
