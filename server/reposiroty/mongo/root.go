package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const MAX_CONNECTION_TIME = 10

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database

	chatCollection *mongo.Collection
}

func NewMongoClient(client *mongo.Client) *MongoClient {
	m := &MongoClient{
		client: client,
	}

	m.db = m.client.Database("streaming-platform")
	m.chatCollection = m.db.Collection("chat-collection")

	return m
}
