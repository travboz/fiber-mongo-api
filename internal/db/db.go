package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInstance struct {
	DbName string
	URI    string
	Client *mongo.Client
}

func NewMongoDBInstance(dbname, uri string) (*MongoDBInstance, error) {
	nm := MongoDBInstance{URI: uri, DbName: dbname}
	err := nm.ConnectToInstance()

	if err != nil {
		return nil, err
	}

	return &nm, nil

}

func (mi *MongoDBInstance) ConnectToInstance() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mi.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB")
	mi.Client = client
	return nil
}

func (mi *MongoDBInstance) GetCollection(collectionName string) *mongo.Collection {
	return mi.Client.Database(mi.DbName).Collection(collectionName)
}
