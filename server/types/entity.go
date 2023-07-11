package types

import "time"

type AuthEntity struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"verified_email"`
	GoogleId  string    `json:"google_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RoomEntity struct {
	RoomHash            string    `json:"room_hash"`
	OwnerName           string    `json:"owner_name"`
	CreatedAt           time.Time `json:"created_at"`
	IsBroadCast         bool      `json:"is_broad_cast"`
	BeforeBroadCastTime int64     `json:"before_broad_cast_time"`
}
