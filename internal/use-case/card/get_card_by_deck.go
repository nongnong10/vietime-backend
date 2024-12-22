package card

import "vietime-backend/internal/entity"

func (uc *cardUseCase) GetCardsByDeck(deckID *string) (*[]entity.Card, error) {
	cards, err := uc.cardRepo.GetCardsByDeck(deckID)
	if err != nil {
		return nil, err
	}
	return cards, nil
}
