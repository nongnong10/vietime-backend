package card

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
	"vietime-backend/internal/repo/card"
)

type CardUseCase interface {
	CreateCard(card *entity.Card) (*entity.Card, error)
	GetCardByID(id *string) (*entity.Card, error)
	UpdateCard(cardID *string, req *dto.UpdateCardRequest) (*entity.Card, error)
	CopyCardToDeck(cardID *string, deckID *string) (*entity.Card, error)
}

type cardUseCase struct {
	cardRepo card.CardRepository
}

func NewCardUseCase(cardRepo card.CardRepository) CardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
	}
}
