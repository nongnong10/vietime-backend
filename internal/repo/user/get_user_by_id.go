package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

func (ur *userRepository) GetByID(id *string) (*entity.User, error) {
	oID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = ur.db.Collection(ur.colName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: oID}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
