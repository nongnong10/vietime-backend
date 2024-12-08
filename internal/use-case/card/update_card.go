package card

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

func (uc *cardUseCase) UpdateCard(cardID *string, req *dto.UpdateCardRequest) (*entity.Card, error) {
	return uc.cardRepo.UpdateCard(cardID, req)
}
