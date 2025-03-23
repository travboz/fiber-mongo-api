package store

import (
	"context"
	"fmt"

	"github.com/travboz/fiber-mongo-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoMongo struct {
	col *mongo.Collection
}

func NewUserRepoMongo(uc *mongo.Collection) *UserRepoMongo {
	return &UserRepoMongo{uc}
}

func (ur *UserRepoMongo) CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	result, err := ur.col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ur *UserRepoMongo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	fmt.Println(objID)
	if err != nil {
		return nil, fmt.Errorf("error converting from hex: %w", err)
	}

	var user models.User

	err = ur.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepoMongo) UpdateUser(ctx context.Context, id string, user models.User) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error converting from hex: %w", err)
	}

	update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}

	result, err := ur.col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

	if err != nil {
		return nil, fmt.Errorf("error updating user in db: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no user found with ID: %s", id)
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := ur.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&updatedUser)

		if err != nil {
			return nil, fmt.Errorf("error updating user in db: %w", err)
		}
	}

	return &updatedUser, nil

}

func (ur *UserRepoMongo) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting from hex: %w", err)
	}

	result, err := ur.col.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("error deleting using: %w", err)
	}

	if result.DeletedCount < 1 {
		return fmt.Errorf("user with specified ID not found")
	}

	return nil
}

func (ur *UserRepoMongo) FetchAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	results, err := ur.col.Find(ctx, bson.M{})

	if err != nil {
		return nil, fmt.Errorf("error finding users: %w", err)
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return nil, fmt.Errorf("error decoding user: %w", err)
		}

		users = append(users, singleUser)
	}

	return users, nil
}
