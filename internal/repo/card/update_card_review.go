package card

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"vietime-backend/internal/entity"
)

func (cr *cardRepo) UpdateCardReview(card *entity.Card) error {
	oID := card.ID
	filter := bson.D{{Key: "_id", Value: oID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "sm2_n", Value: card.Sm2N},
			{Key: "sm2_ef", Value: card.Sm2EF},
			{Key: "sm2_i", Value: card.Sm2I},
			{Key: "last_review", Value: card.LastReview},
			{Key: "next_review", Value: card.NextReview},
			{Key: "num_reviews", Value: card.NumReviews},
		}}}
	_, err := cr.mongodb.Collection(cr.colName).UpdateOne(context.TODO(), &filter, &update)
	if err != nil {
		return err
	}
	return nil
}
