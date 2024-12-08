package card

import (
	"vietime-backend/internal/entity"
)

func (uc *cardUseCase) GetCardByID(id *string) (*entity.Card, error) {
	return uc.cardRepo.GetCardByID(id)
}
