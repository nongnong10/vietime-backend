package card

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

func (cr *cardRepo) GetCardByID(id *string) (*entity.Card, error) {
	oID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	var card entity.Card
	err = cr.mongodb.Collection(cr.colName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: oID}}).Decode(&card)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &card, nil
}
