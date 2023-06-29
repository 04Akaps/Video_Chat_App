package mysql

import (
	"context"
	"database/sql"
	"time"
)

func NewMySqlClient(sourceUri string, option bool) *Queries {
	dbInstance, err := sql.Open("mysql", sourceUri)
	if err != nil {
		panic(err.Error())
	}

	err = dbInstance.Ping()
	if err != nil {
		panic(err.Error())
	}

	if option {
		dbInstance.SetConnMaxLifetime(time.Minute * 1)
		dbInstance.SetMaxIdleConns(3)
		dbInstance.SetMaxOpenConns(6)
	}

	return New(dbInstance)
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}
