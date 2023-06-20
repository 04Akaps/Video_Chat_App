package router

import (
	"github.com/04Akaps/Video_Chat_App/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r *Router) JoinRoom(c *gin.Context) {

	var req types.JoinRoomReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.NewRespHeader(http.StatusBadRequest, "Bind Error : ", err.Error()))
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	r.rooms.IntoRoom(req.RoomId, false, ws)

	for {
		var msg types.BroadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error: ", err)
		}

		msg.Client = ws
		msg.RoomID = req.RoomId

		log.Println(msg.Message)

		r.broadCast <- msg
	}
}

func (r *Router) broadcaster() {
	for {
		msg := <-r.broadCast

		for _, client := range r.rooms.Map[msg.RoomID] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
}
