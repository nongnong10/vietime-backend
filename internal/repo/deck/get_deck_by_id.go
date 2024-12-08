package deck

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

func (dr *deckRepo) GetDeckByID(id *string) (*entity.Deck, error) {
	oID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	var deck entity.Deck
	err = dr.db.Collection(dr.colName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: oID}}).Decode(&deck)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &deck, nil
}
