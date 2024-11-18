package card

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (c *cardRepo) CreateCard(card *entity.Card) (string, error) {
	res, err := c.mongodb.Collection(c.colName).InsertOne(context.TODO(), card)
	if err != nil {
		return "", err
	}

	cardId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("fail to get inserted ID")
	}

	return cardId.Hex(), nil
}
