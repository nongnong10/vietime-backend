package deck

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

func (uc *deckUsecase) UpdateDeck(deckID *string, req *dto.UpdateDeckRequest) (*entity.Deck, error) {
	return uc.deckRepo.UpdateDeck(deckID, req)
}
