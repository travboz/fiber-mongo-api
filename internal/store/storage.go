package store

import (
	"context"

	"github.com/travboz/fiber-mongo-api/internal/db"
	"github.com/travboz/fiber-mongo-api/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	UserRepo UserRepo
}

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

func NewMongoStorage(m *db.MongoDBInstance) Storage {
	return Storage{
		UserRepo: NewUserRepoMongo(m.GetCollection("users")),
	}
}
