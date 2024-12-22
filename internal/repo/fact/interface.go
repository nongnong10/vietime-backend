package fact

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

type FactRepository interface {
	CreateFact(fact *entity.Fact) error
	GetFact() (*entity.Fact, error)
}

type factRepo struct {
	db      *mongo.Database
	colName string
}

func NewFactRepository(mongodb *mongo.Database) FactRepository {
	return &factRepo{
		db:      mongodb,
		colName: "decks",
	}
}
