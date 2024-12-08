package deck

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

func (dr *deckRepo) UpdateDeck(deckID *string, req *dto.UpdateDeckRequest) (*entity.Deck, error) {
	dID, err := primitive.ObjectIDFromHex(*deckID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: dID}}
	update := bson.D{{Key: "$set", Value: *req}}
	option := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedDeck entity.Deck
	err = dr.db.Collection(dr.colName).FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&updatedDeck)
	if err != nil {
		return nil, err
	}
	return &updatedDeck, nil
}
