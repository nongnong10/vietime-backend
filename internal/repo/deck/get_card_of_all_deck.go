package deck

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (dr *deckRepo) GetCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithCards, error) {
	ctx := context.TODO()

	uID, err := primitive.ObjectIDFromHex(*userID)
	if err != nil {
		return nil, err
	}

	// Open an aggregation cursor
	coll := dr.db.Collection(dr.colName)
	cursor, err := coll.Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"user_id", uID}}}},
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
	var deckWithCards []entity.DeckWithCards
	if err = cursor.All(context.TODO(), &deckWithCards); err != nil {
		return nil, err
	}
	return &deckWithCards, nil
}
