package deck

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

type DeckRepository interface {
	CreateDeck(deck *entity.Deck) (*entity.Deck, error)
}

type deckRepo struct {
	db      *mongo.Database
	colName string
}

func NewDeckRepository(mongodb *mongo.Database) DeckRepository {
	return &deckRepo{
		db:      mongodb,
		colName: "decks",
	}
}
