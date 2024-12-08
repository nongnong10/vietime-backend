package deck

import "vietime-backend/internal/entity"

func (uc *deckUsecase) CreateDeck(deck *entity.Deck) (*entity.Deck, error) {
	deck, err := uc.deckRepo.CreateDeck(deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}
