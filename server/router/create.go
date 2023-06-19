package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r *Router) CreateRoom(c *gin.Context) {
	roomID := r.rooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	log.Println(r.rooms.Map)

	c.JSON(http.StatusOK, resp{RoomID: roomID})
}
