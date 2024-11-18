package card

import "vietime-backend/internal/entity"

func (cu *cardUseCase) CreateCard(card *entity.Card) (string, error) {
	return cu.cardRepo.CreateCard(card)
}
