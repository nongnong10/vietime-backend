package fact

import "vietime-backend/internal/entity"

func (fu *factUseCase) GetFact() (*entity.Fact, error) {
	return fu.factRepository.GetFact()
}
