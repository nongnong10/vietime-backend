package card

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (c *cardRepo) CreateCard(card *entity.Card) (*entity.Card, error) {
	card.SetDefault()
	response, err := c.mongodb.Collection(c.colName).InsertOne(context.TODO(), card)
	if err != nil {
		return nil, err
	}

	cardId, ok := response.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("fail to get inserted ID")
	}

	card.ID = cardId

	return card, nil
}
