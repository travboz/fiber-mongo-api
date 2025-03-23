package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInstance struct {
	dbName string
	uri    string
	client *mongo.Client
}

func NewMongoDBInstance(dbname, uri string) *MongoDBInstance {
	return &MongoDBInstance{uri: uri, dbName: dbname}
}

func (mi *MongoDBInstance) ConnectToInstance() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mi.uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB")

	mi.client = client

	return nil
}

// getting database collections
func (mi *MongoDBInstance) GetCollection(collectionName string) *mongo.Collection {
	return mi.client.Database(mi.dbName).Collection(collectionName)
}
