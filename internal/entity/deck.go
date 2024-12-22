package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	time2 "vietime-backend/pkg/utils/time"
)

type Deck struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt           time.Time          `json:"created_at" bson:"created_at"`
	UserID              primitive.ObjectID `json:"user_id" bson:"user_id"`
	IsPublic            bool               `json:"is_public" bson:"is_public"`
	IsFavorite          bool               `json:"is_favorite" bson:"is_favorite"`
	Name                string             `json:"name" bson:"name"`
	Description         string             `json:"description" bson:"description"`
	DescriptionImageURL string             `json:"description_img_url" bson:"description_img_url"`
	Position            string             `json:"position" bson:"position"`
	Views               int                `json:"views" bson:"views"`
	Rating              float32            `json:"rating" bson:"rating"`
	TotalCards          int                `json:"total_cards" bson:"total_cards"`
	TotalLearnedCards   int                `json:"total_learned_cards" bson:"total_learned_cards"`
	MaxNewCards         int                `json:"max_new_cards" bson:"max_new_cards"`
	MaxReviewCards      int                `json:"max_review_cards" bson:"max_review_cards"`
	LastReview          time.Time          `json:"last_review" bson:"last_review"`
	CurNewCards         int                `json:"cur_new_cards" bson:"cur_new_cards"`
	CurReviewCards      int                `json:"cur_review_cards" bson:"cur_review_cards"`
}

type DeckWithReviewCards struct {
	Deck          `bson:"inline"`
	Cards         *[]Card `json:"cards" bson:"cards"`
	NumBlueCards  int     `json:"num_blue_cards" bson:"num_blue_cards"`
	NumRedCards   int     `json:"num_red_cards" bson:"num_red_cards"`
	NumGreenCards int     `json:"num_green_cards" bson:"num_green_cards"`
}

type DeckWithCards struct {
	Deck          `bson:"inline"`
	Cards         *[]Card `json:"cards" bson:"cards"`
	NumBlueCards  int     `json:"num_blue_cards" bson:"num_blue_cards"`
	NumRedCards   int     `json:"num_red_cards" bson:"num_red_cards"`
	NumGreenCards int     `json:"num_green_cards" bson:"num_green_cards"`
}

func (deck *Deck) SetDefault() *Deck {
	deck.CreatedAt = time.Now()
	deck.MaxNewCards = 20
	deck.MaxReviewCards = 100
	deck.Rating = 5
	deck.Views = 1
	deck.TotalLearnedCards = 0
	deck.IsFavorite = false
	return deck
}

func (deck *Deck) UpdateReview() *Deck {
	// Update the LastReview field to the current day
	cur := time2.TruncateToDay(time.Now())
	if cur.Equal(deck.LastReview) {
		return deck
	}
	deck.LastReview = cur
	deck.CurNewCards = 0
	deck.CurReviewCards = 0
	return deck
}
