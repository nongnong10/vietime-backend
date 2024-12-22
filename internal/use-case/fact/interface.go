package fact

import (
	"vietime-backend/internal/entity"
	"vietime-backend/internal/repo/fact"
)

type FactUseCase interface {
	CreateFact(fact *entity.Fact) error
	GetFact() (*entity.Fact, error)
}

type factUseCase struct {
	factRepository fact.FactRepository
}

func NewFactUseCase(factRepository fact.FactRepository) FactUseCase {
	return &factUseCase{
		factRepository: factRepository,
	}
}
