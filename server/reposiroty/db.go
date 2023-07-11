package reposiroty

import (
	"context"
	"database/sql"
	"github.com/04Akaps/Video_Chat_App/config"
	m "github.com/04Akaps/Video_Chat_App/reposiroty/mongo"
	"github.com/04Akaps/Video_Chat_App/reposiroty/mysql"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB struct {
	Mongo *m.MongoClient
	Mysql *mysql.MySqlClient
}

func NewDB(cfg *config.Config) *DB {
	ctx := context.Background()

	db, err := sql.Open(cfg.MySQLConfig.Database, cfg.MySQLConfig.URI)
	if err != nil {
		panic(err)
		return nil
	}

	db.SetMaxIdleConns(cfg.MySQLConfig.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MySQLConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MySQLConfig.ConnMaxLifetime) * time.Second)

	if err = db.Ping(); err != nil {
		panic(err)
		return nil
	}

	mongoConn := options.Client().ApplyURI(cfg.MongoConfig.DatabaseUrl)

	mongoClient, err := mongo.Connect(ctx, mongoConn)

	if err != nil {
		panic(err)
		return nil
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		panic(err)
		return nil
	}

	return &DB{
		Mysql: mysql.NewMySqlClient(db),
		Mongo: m.NewMongoClient(mongoClient),
	}
}
