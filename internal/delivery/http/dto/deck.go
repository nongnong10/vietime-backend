package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

// CreateDeck

type CreateDeckRequest struct {
	UserID              primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	Name                string             `json:"name" binding:"required"`
	Description         string             `json:"description"`
	DescriptionImageURL string             `json:"description_img_url"`
	Position            string             `json:"position"`
	TotalCards          int                `json:"total_cards"`
}

type CreateDeckResponse struct {
	Deck entity.Deck `json:"deck"`
}
