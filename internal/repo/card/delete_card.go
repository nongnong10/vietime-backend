package card

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cr *cardRepo) DeleteCard(cardID *string) error {
	cID, err := primitive.ObjectIDFromHex(*cardID)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: cID}}
	_, err = cr.mongodb.Collection(cr.colName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
