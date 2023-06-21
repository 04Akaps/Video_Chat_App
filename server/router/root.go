package router

import (
	"github.com/04Akaps/Video_Chat_App/config"
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Router struct {
	engine    *gin.Engine
	port      string
	rooms     *RoomMap
	broadCast chan types.BroadcastMsg
	oAuth     *reposiroty.Auth
	paseto    *reposiroty.PasetoMaker
}

func NewRouter(cfg *config.Config) *Router {
	r := Router{
		engine:    gin.New(),
		port:      cfg.ServerInfo.Port,
		broadCast: make(chan types.BroadcastMsg),
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
	r.oAuth = reposiroty.NewOAuth(cfg)
	r.paseto = reposiroty.NewPasetoMaker(cfg.Paseto.PasetoKey)

	r.engine.POST("/create", r.CreateRoom)
	r.engine.POST("/join", r.JoinRoom)

	r.engine.GET("/login", r.login)
	r.engine.GET("/login/callback", r.loginCallback)

	return &r
}

func (r *Router) Run() error {
	log.Println("Http Server Start", "endpoint", r.port)
	return r.engine.Run(r.port)
}
