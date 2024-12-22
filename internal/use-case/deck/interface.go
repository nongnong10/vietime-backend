package deck

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
	cardRepo "vietime-backend/internal/repo/card"
	deckRepo "vietime-backend/internal/repo/deck"
	userRepo "vietime-backend/internal/repo/user"
)

type DeckUsecase interface {
	CreateDeck(deck *entity.Deck) (*entity.Deck, error)
	GetDeckByID(id *string) (*entity.Deck, error)
	UpdateDeck(deckID *string, req *dto.UpdateDeckRequest) (*entity.Deck, error)
	CopyDeck(userID *string, deckID *string) (*entity.DeckWithCards, *entity.DeckWithReviewCards, error)
}

type deckUsecase struct {
	deckRepo deckRepo.DeckRepository
	userRepo userRepo.UserRepository
	cardRepo cardRepo.CardRepository
}

func NewDeckUsecase(
	deckRepo deckRepo.DeckRepository,
	userRepo userRepo.UserRepository,
	cardRepo cardRepo.CardRepository,
) DeckUsecase {
	return &deckUsecase{
		deckRepo: deckRepo,
		userRepo: userRepo,
		cardRepo: cardRepo,
	}
}
