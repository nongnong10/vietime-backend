package deck

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/helpers"
)

func (uc *deckUsecase) GetReviewCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error) {
	rawDeckWithCards, err := uc.deckRepo.GetCardsAllDecksOfUser(userID)
	if err != nil {
		return nil, err
	}
	decksWithReviewCards := []entity.DeckWithReviewCards{}
	for i := range *rawDeckWithCards {
		rawDeck := (*rawDeckWithCards)[i]
		deck := entity.DeckWithReviewCards{
			Deck:          rawDeck.Deck, // Copy fields from the embedded Deck
			Cards:         rawDeck.Cards,
			NumBlueCards:  0, // Set the initial values for NumBlueCards, NumRedCards, NumGreenCards
			NumRedCards:   0,
			NumGreenCards: 0,
		}
		deck.UpdateReview()
		deck.Cards, deck.NumBlueCards, deck.NumRedCards, deck.NumGreenCards = helpers.FilterReviewCards(deck.Cards, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
		decksWithReviewCards = append(decksWithReviewCards, deck)
	}
	return &decksWithReviewCards, nil
}
