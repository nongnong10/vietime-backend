package card

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

func (uc *cardUseCase) CopyCardToDeck(cardID *string, deckID *string) (*entity.Card, error) {
	card, err := uc.cardRepo.GetCardByID(cardID)
	if err != nil {
		return nil, err
	}
	card.ID = primitive.NilObjectID
	card.DeckID, err = primitive.ObjectIDFromHex(*deckID)
	card.SetDefault()
	if err != nil {
		return nil, err
	}
	card, err = uc.cardRepo.CreateCard(card)
	if err != nil {
		return nil, err
	}
	return card, nil
}
