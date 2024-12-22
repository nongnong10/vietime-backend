package card

import "vietime-backend/internal/entity"

func (uc *cardUseCase) UpdateCardReview(card *entity.Card) error {
	return uc.cardRepo.UpdateCardReview(card)
}
