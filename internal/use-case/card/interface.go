package card

import (
	"vietime-backend/internal/entity"
	"vietime-backend/internal/repo/card"
)

type CardUseCase interface {
	CreateCard(card *entity.Card) (*entity.Card, error)
}

type cardUseCase struct {
	cardRepo card.CardRepository
}

func NewCardUseCase(cardRepo card.CardRepository) CardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
	}
}
