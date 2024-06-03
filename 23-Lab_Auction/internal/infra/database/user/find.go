package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/logger"
	user_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/user"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *UserRepository) FindUserById(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var userEntityMongo UserEntityMongo

	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", id)
			logger.Error(errorMessage, err)
			return nil, internal_error.NewNotFoundError(errorMessage)
		}

		logger.Error("Error trying to find user by id", err)
		return nil, internal_error.NewNotFoundError("Error trying to find user by id")
	}

	return &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}, nil
}
