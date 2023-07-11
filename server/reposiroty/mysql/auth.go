package mysql

import (
	"context"
	"github.com/04Akaps/Video_Chat_App/types"
)

type Auth struct{}

func (auth *Auth) Insert(ctx context.Context, db ISqlContext, name, verifiedEmail, googleId string) error {
	_, err := db.ExecContext(ctx, "INSERT INTO chat.auth(name, verified_email, google_id) VALUES(?, ?, ?);", name, verifiedEmail, googleId)
	return err
}

func (auth *Auth) GetAuthByName(ctx context.Context, db ISqlContext, name string) (*types.AuthEntity, error) {
	res := new(types.AuthEntity)

	err := db.QueryRowContext(ctx, "SELECT * FROM chat.auth WHERE name = ?;", name).Scan(
		&res.Id,
		&res.Name,
		&res.Email,
		&res.GoogleId,
		&res.CreatedAt,
	)

	return res, err
}
