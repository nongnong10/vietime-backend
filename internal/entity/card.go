package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Card struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`
	DeckID         primitive.ObjectID `json:"deck_id" bson:"deck_id"`
	Index          int32              `json:"index" bson:"index"`
	Question       string             `json:"question" bson:"question"`
	QuestionImgUrl string             `json:"question_img_url" bson:"question_img_url"`
	Answer         string             `json:"answer" bson:"answer"`
	WrongAnswers   []string           `json:"wrong_answers" bson:"wrong_answers"`
	LastReview     time.Time          `json:"last_review" bson:"last_review"`
	NextReview     time.Time          `json:"next_review" bson:"next_review"`
	NumReviews     int32              `json:"num_reviews" bson:"num_reviews"`
	CardType       string             `json:"card_type" bson:"card_type"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
}

func (card *Card) SetDefault() *Card {
	card.CreatedAt = time.Now()
	card.LastReview = time.Now()
	card.NextReview = time.Now()
	card.NumReviews = 0
	return card
}
