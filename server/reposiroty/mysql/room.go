package mysql

import (
	"context"
	"github.com/04Akaps/Video_Chat_App/types"
)

type Room struct{}

func (room *Room) Insert(ctx context.Context, db ISqlContext, hash, ownerName string) error {
	_, err := db.ExecContext(ctx, "INSERT INTO chat.room(owner_name, room_hash) VALUES(?, ?);", ownerName, hash)
	return err
}

func (room *Room) GetRoomByHash(ctx context.Context, db ISqlContext, hash string) (*types.RoomEntity, error) {
	res := new(types.RoomEntity)

	err := db.QueryRowContext(ctx, "SELECT * FROM chat.room WHERE room_hash = ?;", hash).Scan(
		&res.RoomHash,
		&res.OwnerName,
		&res.CreatedAt,
	)

	return res, err
}

func (room *Room) GetRoomByOwnerName(ctx context.Context, db ISqlContext, owner string) ([]*types.RoomEntity, error) {

	if row, err := db.QueryContext(ctx, "SELECT * FROM chat.room WHERE owner_name = ?;", owner); err != nil {
		return nil, err
	} else {
		defer row.Close()

		res := make([]*types.RoomEntity, 0)

		for row.Next() {
			model := new(types.RoomEntity)

			if err := row.Scan(
				&model.RoomHash,
				&model.OwnerName,
				&model.CreatedAt,
			); err != nil {
				return nil, err
			} else {
				res = append(res, model)
			}
		}

		return res, nil
	}
}

func (room *Room) GetAllRoomByPaging(ctx context.Context, db ISqlContext, start, end int64) ([]*types.RoomEntity, error) {
	if row, err := db.QueryContext(ctx, "SELECT * FROM chat.room LIMIT ?, ?", start, end); err != nil {
		return nil, err
	} else {
		defer row.Close()

		res := make([]*types.RoomEntity, 0)

		for row.Next() {
			model := new(types.RoomEntity)

			if err := row.Scan(
				&model.RoomHash,
				&model.OwnerName,
				&model.CreatedAt,
			); err != nil {
				return nil, err
			} else {
				res = append(res, model)
			}
		}

		return res, nil
	}
}
