package mysql

import (
	"context"
	"github.com/04Akaps/Video_Chat_App/types"
	"time"
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
		&res.IsBroadCast,
		&res.BeforeBroadCastTime,
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
				&model.IsBroadCast,
				&model.BeforeBroadCastTime,
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

func (room *Room) GetRoomCountByName(ctx context.Context, db ISqlContext, owner string) (int64, error) {

	type count struct {
		TotalCount int64 `json:"total_count"`
	}

	var queryResult count

	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) AS total_count FROM chat.room WHERE owner_name = ?;", owner).Scan(&queryResult.TotalCount); err != nil {
		return 0, err
	} else {
		return queryResult.TotalCount, nil
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
				&model.IsBroadCast,
				&model.BeforeBroadCastTime,
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

func (room *Room) UpdateBroadCast(ctx context.Context, db ISqlContext, updateStatus bool, hash string) error {
	_, err := db.ExecContext(ctx, "UPDATE chat.room SET (is_broad_cast = ?, before_broad_cast_time = ?) WHERE room_hash = ?", updateStatus, time.Now(), hash)
	return err
}

func (room *Room) RecentlyCreatedRoomLIst(ctx context.Context, db ISqlContext) ([]*types.RoomEntity, error) {

	if row, err := db.QueryContext(ctx, "SELECT * FROM chat.room ORDER BY created_at DESC LIMIT 5;"); err != nil {
		return nil, err
	} else {
		defer row.Close()

		res := make([]*types.RoomEntity, 0)

		for row.Next() {
			model := new(types.RoomEntity)

			if err := row.Scan(
				&model.RoomHash,
				&model.OwnerName,
				&model.IsBroadCast,
				&model.BeforeBroadCastTime,
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
