package deck

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

type DeckRepository interface {
	CreateDeck(deck *entity.Deck) (*entity.Deck, error)
	GetDeckByID(id *string) (*entity.Deck, error)
	GetDeckWithCards(deckID *string) (*entity.DeckWithCards, error)
	UpdateDeck(deckID *string, req *dto.UpdateDeckRequest) (*entity.Deck, error)
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
