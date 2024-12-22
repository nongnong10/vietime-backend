package deck

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/helpers"
)

func (uc *deckUsecase) GetDecksWithCards(userID *string) (*[]entity.DeckWithCards, *[]entity.DeckWithCards, *[]entity.DeckWithReviewCards, error) {
	rawDeckWithCards, err := uc.deckRepo.GetCardsAllDecks(userID)
	if err != nil {
		return nil, nil, nil, err
	}
	userDecks := []entity.DeckWithCards{}
	publicDecks := []entity.DeckWithCards{}
	for _, deck := range *rawDeckWithCards {
		if deck.IsPublic {
			publicDecks = append(publicDecks, deck)
		} else {
			userDecks = append(userDecks, deck)
		}
	}
	decksWithReviewCards := []entity.DeckWithReviewCards{}
	for i := range userDecks {
		rawDeck := userDecks[i]
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
		userDecks[i].NumBlueCards = deck.NumBlueCards
		userDecks[i].NumRedCards = deck.NumRedCards
		userDecks[i].NumGreenCards = deck.NumGreenCards
	}
	return &userDecks, &publicDecks, &decksWithReviewCards, nil
}
