package card

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/helpers"
)

func (uc *cardUseCase) GetReviewCardsByDeck(deckID *string, maxNewCards int, maxReviewCards int) (*[]entity.Card, int, int, int, error) {
	cards, err := uc.cardRepo.GetCardsByDeck(deckID)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	cards, numBlue, numRed, numGreen := helpers.FilterReviewCards(cards, maxNewCards, maxReviewCards)
	return cards, numBlue, numRed, numGreen, nil
}
