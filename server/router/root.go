package router

import (
	"errors"
	"github.com/04Akaps/Video_Chat_App/config"
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type Router struct {
	engine    *gin.Engine
	port      string
	rooms     *reposiroty.RoomMap
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

	r.rooms = reposiroty.NewAllRooms()
	r.oAuth = reposiroty.NewOAuth(cfg)
	r.paseto = reposiroty.NewPasetoMaker(cfg.Paseto.PasetoKey)

	newAuth(r, r.paseto)
	newRoom(r, r.rooms, r.broadCast)

	return &r
}

func (r *Router) Run() error {
	log.Println("Http Server Start", "endpoint", r.port)
	return r.engine.Run(r.port)
}

func (p *Router) extractToken(c *gin.Context) error {
	bearToken := c.Request.Header.Get("Authorization")
	slice := strings.Split(bearToken, " ") //"bearer" 제거

	if len(slice) < 1 {
		return errors.New("Not Found Auth Token")
	}

	return p.paseto.VerifyToken(slice[1])
}

func (r *Router) verifyAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := r.extractToken(c); err != nil {
			c.JSON(http.StatusForbidden, "Failed Auth")
		} else {
			c.Next() //문제가 없다면 다음 handler 호출.
		}
	}
}
