package deck

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/helpers"
)

func (uc *deckUsecase) CopyDeck(userID *string, deckID *string) (*entity.DeckWithCards, *entity.DeckWithReviewCards, error) {
	// 1. Get info of deck
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, nil, err
	}
	deck, err := uc.deckRepo.GetDeckByID(deckID)
	if err != nil {
		return nil, nil, err
	}
	cards, err := uc.cardRepo.GetCardsByDeck(deckID)
	if err != nil {
		return nil, nil, err
	}

	deck.ID = primitive.NilObjectID
	deck.UserID = user.ID
	deck.IsPublic = false

	// 2. Clone new deck
	deck, err = uc.deckRepo.CreateDeck(deck)
	if err != nil {
		return nil, nil, err
	}

	// 3. Create new cards with the cards of that deck
	for i := range *cards {
		(*cards)[i].UserID = deck.UserID
		(*cards)[i].DeckID = deck.ID
		(*cards)[i].ID = primitive.NilObjectID
	}
	err = uc.cardRepo.CreateManyCards(*cards)
	if err != nil {
		return nil, nil, err
	}

	// 4. Update review of deck
	dID := deck.ID.Hex()
	rawDeckWithCards, err := uc.deckRepo.GetDeckWithCards(&dID)
	if err != nil {
		return nil, nil, err
	}
	deckWithReviewCards := entity.DeckWithReviewCards{
		Deck:          rawDeckWithCards.Deck,
		Cards:         rawDeckWithCards.Cards,
		NumBlueCards:  0,
		NumRedCards:   0,
		NumGreenCards: 0,
	}
	deckWithReviewCards.UpdateReview()
	deckWithReviewCards.Cards, deckWithReviewCards.NumBlueCards, deckWithReviewCards.NumRedCards, deckWithReviewCards.NumGreenCards =
		helpers.FilterReviewCards(
			deckWithReviewCards.Cards,
			deckWithReviewCards.MaxNewCards-deckWithReviewCards.CurNewCards,
			deckWithReviewCards.MaxReviewCards-deckWithReviewCards.CurReviewCards,
		)
	rawDeckWithCards.NumBlueCards = deckWithReviewCards.NumBlueCards
	rawDeckWithCards.NumRedCards = deckWithReviewCards.NumRedCards
	rawDeckWithCards.NumGreenCards = deckWithReviewCards.NumGreenCards
	return rawDeckWithCards, &deckWithReviewCards, nil
}
