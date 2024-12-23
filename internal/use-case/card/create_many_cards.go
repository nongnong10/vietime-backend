package card

import "vietime-backend/internal/entity"

func (cu *cardUseCase) CreateManyCards(cards []entity.Card) error {
	return cu.cardRepo.CreateManyCards(cards)
}
