package mysql

import (
	"context"
	"github.com/04Akaps/Video_Chat_App/types"
)

type Participant struct{}

func (participant *Participant) GetParticipant(ctx context.Context, db ISqlContext, hash string) ([]*types.RoomParticipantEntity, error) {

	if row, err := db.QueryContext(ctx, "SELECT user_name FROM chat.room_participant WHERE room_hash = ?;", hash); err != nil {
		return nil, err
	} else {
		defer row.Close()

		res := make([]*types.RoomParticipantEntity, 0)

		for row.Next() {
			model := new(types.RoomParticipantEntity)

			if err := row.Scan(
				&model.UserName,
			); err != nil {
				return nil, err
			} else {
				res = append(res, model)
			}
		}

		return res, nil
	}
}
