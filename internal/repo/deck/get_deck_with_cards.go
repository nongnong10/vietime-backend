package deck

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

func (dr *deckRepo) GetDeckWithCards(deckID *string) (*entity.DeckWithCards, error) {
	ctx := context.TODO()

	dID, err := primitive.ObjectIDFromHex(*deckID)
	if err != nil {
		return nil, err
	}

	// Open an aggregation cursor
	coll := dr.db.Collection(dr.colName)
	cursor, err := coll.Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"_id", dID}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "cards"},
					{"localField", "_id"},
					{"foreignField", "deck_id"},
					{"as", "cards"},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"cards",
						bson.D{
							{"$sortArray",
								bson.D{
									{"input", "$cards"},
									{"sortBy", bson.D{{"created_at", 1}}},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var deckWithCardsList []entity.DeckWithCards
	if err = cursor.All(context.TODO(), &deckWithCardsList); err != nil {
		return nil, err
	}

	// Take the first element if available
	if len(deckWithCardsList) > 0 {
		return &deckWithCardsList[0], nil
	}

	// Handle case where no documents are found
	return nil, mongo.ErrNoDocuments
}
