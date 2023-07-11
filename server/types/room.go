package types

type Auth struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"verified_email"`
	GoogleId  string `json:"google_id"`
	CreatedAt int64  `json:"created_at"`
}

type Room struct {
	RoomHash            string `json:"room_hash"`
	OwnerName           string `json:"owner_name"`
	CreatedAt           int64  `json:"created_at"`
	IsBroadCast         bool   `json:"is_broad_cast"`
	BeforeBroadCastTime int64  `json:"before_broad_cast_time"`
}

type TotalRoomListReq struct {
	Paging
}

type GetRoomByHashReq struct {
	Hash string `uri:"hash"`
}

type GetRoomByHashRes struct {
	Room *Room `json:"room"`
}

type ChangeRoomStatusByHashReq struct {
	Hash string `uri:"hash"`
}
