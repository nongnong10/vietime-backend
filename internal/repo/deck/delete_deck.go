package deck

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dr *deckRepo) DeleteDeck(deckID *string) error {
	dID, err := primitive.ObjectIDFromHex(*deckID)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: dID}}
	_, err = dr.db.Collection(dr.colName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	filter = bson.D{{Key: "deck_id", Value: dID}}
	_, err = dr.db.Collection("cards").DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
