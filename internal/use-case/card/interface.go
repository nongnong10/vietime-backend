package card

import (
	"vietime-backend/internal/entity"
	"vietime-backend/internal/repo/card"
)

type CardUseCase interface {
	CreateCard(card *entity.Card) (*entity.Card, error)
}

type cardUseCase struct {
	cardRepo card.CardRepo
}

func NewCardUseCase(cardRepo card.CardRepo) CardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
	}
}
