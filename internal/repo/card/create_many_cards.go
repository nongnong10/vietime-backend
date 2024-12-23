package card

import (
	"context"
	"vietime-backend/internal/entity"
)

func (cr *cardRepo) CreateManyCards(cards []entity.Card) error {
	for i := range cards {
		cards[i].SetDefault()
	}
	newCards := make([]interface{}, len(cards))
	for i := range cards {
		newCards[i] = cards[i]
	}
	_, err := cr.mongodb.Collection(cr.colName).InsertMany(context.TODO(), newCards)
	if err != nil {
		return err
	}
	return nil
}
