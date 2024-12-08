package deck

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
	deckRepo "vietime-backend/internal/repo/deck"
	userRepo "vietime-backend/internal/repo/user"
)

type DeckUsecase interface {
	CreateDeck(deck *entity.Deck) (*entity.Deck, error)
	GetDeckByID(id *string) (*entity.Deck, error)
	UpdateDeck(deckID *string, req *dto.UpdateDeckRequest) (*entity.Deck, error)
}

type deckUsecase struct {
	deckRepo deckRepo.DeckRepository
	userRepo userRepo.UserRepository
}

func NewDeckUsecase(deckRepo deckRepo.DeckRepository, userRepo userRepo.UserRepository) DeckUsecase {
	return &deckUsecase{
		deckRepo: deckRepo,
		userRepo: userRepo,
	}
}
