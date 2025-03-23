package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")

	return client, nil
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golang-api").Collection(collectionName)
	return collection
}
