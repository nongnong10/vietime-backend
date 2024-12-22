package card

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

type CardRepository interface {
	CreateCard(card *entity.Card) (*entity.Card, error)
	CreateManyCards(cards *[]entity.Card) error
	GetCardByID(id *string) (*entity.Card, error)
	GetCardsByDeck(deckID *string) (*[]entity.Card, error)
	UpdateCard(cardID *string, req *dto.UpdateCardRequest) (*entity.Card, error)
	UpdateCardReview(card *entity.Card) error
}

type cardRepo struct {
	mongodb *mongo.Database
	colName string
}

func NewCardRepository(mongodb *mongo.Database) CardRepository {
	return &cardRepo{
		mongodb: mongodb,
		colName: "cards",
	}
}
