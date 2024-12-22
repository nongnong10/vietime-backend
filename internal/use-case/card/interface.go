package card

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
	"vietime-backend/internal/repo/card"
)

type CardUseCase interface {
	CreateCard(card *entity.Card) (*entity.Card, error)

	GetCardByID(id *string) (*entity.Card, error)
	GetCardsByDeck(deckID *string) (*[]entity.Card, error)
	GetReviewCardsByDeck(deckID *string, maxNewCards int, maxReviewCards int) (*[]entity.Card, int, int, int, error)

	UpdateCard(cardID *string, req *dto.UpdateCardRequest) (*entity.Card, error)
	UpdateCardReview(card *entity.Card) error

	CopyCardToDeck(cardID *string, deckID *string) (*entity.Card, error)

	DeleteCard(cardID *string) error
}

type cardUseCase struct {
	cardRepo card.CardRepository
}

func NewCardUseCase(cardRepo card.CardRepository) CardUseCase {
	return &cardUseCase{
		cardRepo: cardRepo,
	}
}
