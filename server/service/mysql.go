package service

import (
	"context"
	"errors"
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/types"
	"log"
)

type mysql struct {
	db *reposiroty.DB
}

func newMySqlService(db *reposiroty.DB) *mysql {
	return &mysql{
		db: db,
	}
}

func (m *mysql) GetAuthByName(userName string) (*types.Auth, error) {
	if res, err := m.db.Mysql.Auth.GetAuthByName(context.TODO(), m.db.Mysql.DB, userName); err != nil {
		return nil, err
	} else {
		return &types.Auth{
			Id:        res.Id,
			Name:      res.Name,
			Email:     res.Email,
			GoogleId:  res.GoogleId,
			CreatedAt: res.CreatedAt.Unix(),
		}, nil
	}
}

func (m *mysql) InsertAuth(name, verifiedEmail, googleId string) error {
	return m.db.Mysql.Auth.Insert(context.TODO(), m.db.Mysql.DB, name, verifiedEmail, googleId)
}

func (m *mysql) InsertRoom(hash, ownerName string) error {
	return m.db.Mysql.Room.Insert(context.TODO(), m.db.Mysql.DB, hash, ownerName)
}

func (m *mysql) GetRoomDataByHash(hash string) (*types.Room, error) {
	if tx, err := m.db.Mysql.DB.Begin(); err != nil {
		return nil, err
	} else {
		room, err := m.db.Mysql.Room.GetRoomByHash(context.TODO(), tx, hash)

		if err != nil {
			if e := tx.Rollback(); e != nil {
				// TODO  - 로깅 넣읍시다.!!
				log.Println("RollBack : ", e)
			}
			return nil, err
		}

		return &types.Room{
			RoomHash:            room.RoomHash,
			OwnerName:           room.OwnerName,
			CreatedAt:           room.CreatedAt.Unix(),
			IsBroadCast:         room.IsBroadCast,
			BeforeBroadCastTime: room.BeforeBroadCastTime,
		}, tx.Commit()
	}

}

func (m *mysql) GetRoomByOwner(owner string) ([]*types.Room, error) {
	if res, err := m.db.Mysql.Room.GetRoomByOwnerName(context.TODO(), m.db.Mysql.DB, owner); err != nil {
		return nil, err
	} else {

		if len(res) == 0 {
			return []*types.Room{}, nil
		}
		var model []*types.Room

		for _, room := range res {
			newModel := &types.Room{
				RoomHash:            room.RoomHash,
				OwnerName:           room.OwnerName,
				CreatedAt:           room.CreatedAt.Unix(),
				IsBroadCast:         room.IsBroadCast,
				BeforeBroadCastTime: room.BeforeBroadCastTime,
			}

			model = append(model, newModel)
		}

		return model, nil
	}
}

func (m *mysql) RoomCountByName(owner string) (int64, error) {
	return m.db.Mysql.Room.GetRoomCountByName(context.TODO(), m.db.Mysql.DB, owner)
}
func (m *mysql) GetAllRoomByPaging(paging *types.Paging) ([]*types.Room, error) {
	verifyPagingOption(paging)

	if res, err := m.db.Mysql.Room.GetAllRoomByPaging(context.TODO(), m.db.Mysql.DB, paging.Page, paging.PageSize); err != nil {
		return nil, err
	} else {
		var model []*types.Room

		for _, room := range res {
			newModel := &types.Room{
				RoomHash:            room.RoomHash,
				OwnerName:           room.OwnerName,
				CreatedAt:           room.CreatedAt.Unix(),
				IsBroadCast:         room.IsBroadCast,
				BeforeBroadCastTime: room.BeforeBroadCastTime,
			}

			model = append(model, newModel)
		}

		return model, nil
	}
}

func (m *mysql) ChangeBroadCastStatus(owner, hash string) error {
	if result, err := m.GetRoomDataByHash(hash); err != nil {
		return err
	} else {
		if result.OwnerName != owner {
			return errors.New("Not Your Room")
		} else {
			return m.db.Mysql.Room.UpdateBroadCast(context.TODO(), m.db.Mysql.DB, !result.IsBroadCast, hash)
		}
	}
}

func (m *mysql) RecentlyCreatedRoomList() ([]*types.Room, error) {
	if res, err := m.db.Mysql.Room.RecentlyCreatedRoomLIst(context.TODO(), m.db.Mysql.DB); err != nil {
		return nil, err
	} else {

		if len(res) == 0 {
			return []*types.Room{}, nil
		}
		var model []*types.Room

		for _, room := range res {
			newModel := &types.Room{
				RoomHash:            room.RoomHash,
				OwnerName:           room.OwnerName,
				CreatedAt:           room.CreatedAt.Unix(),
				IsBroadCast:         room.IsBroadCast,
				BeforeBroadCastTime: room.BeforeBroadCastTime,
			}

			model = append(model, newModel)
		}

		return model, nil
	}
}
