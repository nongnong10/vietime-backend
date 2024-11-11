package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (ur *userRepository) Create(user *entity.User) (string, error) {
	user.SetDefault()
	result, err := ur.db.Collection(ur.colName).InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted ID")
	}

	return insertedID.Hex(), nil
}
