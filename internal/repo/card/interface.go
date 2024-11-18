package card

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

type CardRepo interface {
	CreateCard(card *entity.Card) (string, error)
}

type cardRepo struct {
	mongodb *mongo.Database
	colName string
}

func NewCardRepo(mongodb *mongo.Database) CardRepo {
	return &cardRepo{
		mongodb: mongodb,
		colName: "cards",
	}
}
