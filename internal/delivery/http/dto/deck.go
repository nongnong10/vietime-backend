package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

// Update Deck
type UpdateDeckRequest struct {
	DeckID              *primitive.ObjectID `json:"deck_id" bson:"_id,omitempty" binding:"required"`
	IsPublic            *bool               `json:"is_public" bson:"is_public,omitempty"`
	IsFavorite          *bool               `json:"is_favorite" bson:"is_favorite,omitempty"`
	Name                *string             `json:"name" bson:"name,omitempty"`
	Description         *string             `json:"description" bson:"description,omitempty"`
	DescriptionImageURL *string             `json:"description_img_url" bson:"description_img_url,omitempty"`
	Position            *string             `json:"position" bson:"position,omitempty"`
	TotalLearnedCards   *int                `json:"total_learned_cards" bson:"total_learned_cards,omitempty"`
	MaxNewCards         *int                `json:"max_new_cards" bson:"max_new_cards,omitempty"`
	MaxReviewCards      *int                `json:"max_review_cards" bson:"max_review_cards,omitempty"`
	LastReview          *time.Time          `json:"last_review" bson:"last_review,omitempty"`
	CurNewCards         *int                `json:"cur_new_cards" bson:"cur_new_cards,omitempty"`
	CurReviewCards      *int                `json:"cur_review_cards" bson:"cur_review_cards,omitempty"`
	Views               *int                `json:"views" bson:"views,omitempty"`
	Rating              *float32            `json:"rating" bson:"rating,omitempty"`
}

type UpdateDeckResponse struct {
	Success bool        `json:"success"`
	Deck    entity.Deck `json:"deck"`
}

type UpdateViewDeckRequest struct {
	DeckID *primitive.ObjectID `json:"deck_id" bson:"_id,omitempty" binding:"required"`
	Views  *int                `json:"views" bson:"views,omitempty"`
}

// Copy Deck
type CopyDeckRequest struct {
	DeckID *primitive.ObjectID `json:"deck_id" binding:"required"`
}

type CopyDeckResponse struct {
	Deck       entity.DeckWithCards       `json:"deck"`
	DeckReview entity.DeckWithReviewCards `json:"deck_review"`
}

// Delete Deck
type DeleteDeckRequest struct {
	DeckID *primitive.ObjectID `json:"deck_id" binding:"required"`
}

type DeleteDeckResponse struct {
	Success bool `json:"success"`
}

// Success Response
type SuccessResponse struct {
	Success bool `json:"success"`
}
