package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

func (ur *userRepository) GetByEmail(email *string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Collection(ur.colName).FindOne(context.TODO(), bson.D{{Key: "email", Value: *email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
