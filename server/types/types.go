package types

type JoinRoomReq struct {
	RoomId int64 `form:"roomId" binding:"required"`
}
