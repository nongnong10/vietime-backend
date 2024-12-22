package card

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (cr *cardRepo) GetCardsByDeck(deckID *string) (*[]entity.Card, error) {
	dID, err := primitive.ObjectIDFromHex(*deckID)
	if err != nil {
		return nil, err
	}
	cursor, err := cr.mongodb.Collection(cr.colName).Find(
		context.TODO(),
		bson.D{{Key: "deck_id", Value: dID}},
	)
	if err != nil {
		return nil, err
	}

	var cards []entity.Card
	if err = cursor.All(context.TODO(), &cards); err != nil {
		return nil, err
	}

	return &cards, nil
}
