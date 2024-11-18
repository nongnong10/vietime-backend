package card

import "vietime-backend/internal/repo/card"

type CardUseCase interface {
}

type cardUseCase struct {
	cardRepo card.CardRepo
}

func NewCardUseCase(cardRepo card.CardRepo) CardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
	}
}
