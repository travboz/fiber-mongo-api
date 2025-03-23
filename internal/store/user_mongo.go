package store

import (
	"context"

	"github.com/travboz/fiber-mongo-api/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoMongo struct {
	UsersCollection *mongo.Collection
}

func NewUserRepoMongo(uc *mongo.Collection) UserRepo {
	return &UserRepoMongo{uc}
}

func (ur *UserRepoMongo) CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	return nil, nil
}
func (ur *UserRepoMongo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
