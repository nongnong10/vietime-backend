package deck

import "vietime-backend/internal/entity"

func (uc *deckUsecase) GetDeckByID(id *string) (*entity.Deck, error) {
	return uc.deckRepo.GetDeckByID(id)
}
