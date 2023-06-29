package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

const MAX_CONNECTION_TIME = 10

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database

	chatCollection *mongo.Collection
}

func NewMongoClient(sourceUri string, option bool) *MongoClient {
	ctx := context.Background()

	m := &MongoClient{}

	var client *mongo.Client
	var err error

	mongoConn := options.Client().ApplyURI(sourceUri)

	if option {
		jmajority := writeconcern.New(writeconcern.J(true))
		wmajority := writeconcern.New(writeconcern.W(1))
		tmajority := writeconcern.New(writeconcern.WTimeout(1000 * time.Microsecond))
		readConcert := readconcern.New(readconcern.Level("majority"))

		mongoConn.
			SetConnectTimeout(MAX_CONNECTION_TIME * time.Second).
			SetMaxPoolSize(50).SetMinPoolSize(5).
			SetWriteConcern(jmajority).
			SetWriteConcern(wmajority).
			SetWriteConcern(tmajority).
			SetReadConcern(readConcert)
	}

	if m.client, err = mongo.Connect(ctx, mongoConn); err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	m.db = m.client.Database("streaming-platform")
	m.chatCollection = m.db.Collection("chat-collection")

	return m
}
