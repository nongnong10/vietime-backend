package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"vietime-backend/internal/delivery/http/request"
	"vietime-backend/internal/entity"
)

func (ur *userRepository) UpdateUser(userID *string, req *request.UpdateUserRequest) (*entity.User, error) {
	uID, err := primitive.ObjectIDFromHex(*userID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: uID}}
	update := bson.D{{Key: "$set", Value: *req}}
	option := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedUser entity.User
	err = ur.db.Collection(ur.colName).FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}
