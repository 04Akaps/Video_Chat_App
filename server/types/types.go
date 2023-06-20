package types

import "github.com/gorilla/websocket"

type JoinRoomReq struct {
	RoomId string `form:"roomId" binding:"required"`
}

type BroadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}
