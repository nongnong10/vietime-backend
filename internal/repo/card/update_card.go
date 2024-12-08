package card

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

func (cr *cardRepo) UpdateCard(cardID *string, req *dto.UpdateCardRequest) (*entity.Card, error) {
	cID, err := primitive.ObjectIDFromHex(*cardID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: cID}}
	update := bson.D{{Key: "$set", Value: *req}}
	option := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedCard entity.Card
	err = cr.mongodb.Collection(cr.colName).FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&updatedCard)
	if err != nil {
		return nil, err
	}
	return &updatedCard, nil
}
