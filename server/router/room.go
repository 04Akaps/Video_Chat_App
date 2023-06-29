package router

import (
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type room struct {
	router    Router
	rooms     *reposiroty.RoomMap
	broadCast chan types.BroadcastMsg
	upgrader  websocket.Upgrader
}

func newRoom(r Router, rooms *reposiroty.RoomMap, broadCast chan types.BroadcastMsg) *room {
	a := &room{
		router:    r,
		rooms:     rooms,
		broadCast: broadCast,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	baseUri := "/room"

	r.engine.POST(baseUri+"/create", a.router.verifyAuthToken(), a.CreateRoom)
	r.engine.POST(baseUri+"/join", a.router.verifyAuthToken(), a.JoinRoom)

	return a
}

func (r *room) CreateRoom(c *gin.Context) {
	roomID := r.rooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	c.JSON(http.StatusOK, resp{RoomID: roomID})
}

func (r *room) JoinRoom(c *gin.Context) {

	var req types.JoinRoomReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.NewRespHeader(http.StatusBadRequest, "Bind Error : ", err.Error()))
		return
	}

	ws, err := r.upgrader.Upgrade(c.Writer, c.Request, nil)

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

func (r *room) broadcaster() {
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
