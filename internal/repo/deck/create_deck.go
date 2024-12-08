package deck

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (dr *deckRepo) CreateDeck(deck *entity.Deck) (*entity.Deck, error) {
	deck.SetDefault()
	result, err := dr.db.Collection(dr.colName).InsertOne(context.TODO(), deck)
	if err != nil {
		return nil, err
	}
	deck.ID = result.InsertedID.(primitive.ObjectID)
	return deck, nil
}
