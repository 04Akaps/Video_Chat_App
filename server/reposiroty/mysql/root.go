package mysql

import (
	"context"
	"database/sql"
)

type ISqlContext interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type MySqlClient struct {
	DB   *sql.DB
	Auth *Auth
	Room *Room
}

func NewMySqlClient(db *sql.DB) *MySqlClient {
	return &MySqlClient{
		DB:   db,
		Auth: &Auth{},
		Room: &Room{},
	}
}
