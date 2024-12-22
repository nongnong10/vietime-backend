package fact

import "vietime-backend/internal/entity"

func (fu *factUseCase) CreateFact(fact *entity.Fact) error {
	return fu.factRepository.CreateFact(fact)
}
