package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Card struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	DeckID           primitive.ObjectID `json:"deck_id" bson:"deck_id"`
	Index            int                `json:"index" bson:"index"`
	Question         string             `json:"question" bson:"question"`
	QuestionImgURL   string             `json:"question_img_url" bson:"question_img_url"`
	QuestionImgLabel string             `json:"question_img_label" bson:"question_img_label"`
	Answer           string             `json:"answer" bson:"answer"`
	WrongAnswers     []string           `json:"wrong_answers" bson:"wrong_answers"`
	LastReview       time.Time          `json:"last_review" bson:"last_review"`
	NextReview       time.Time          `json:"next_review" bson:"next_review"`
	NumReviews       int                `json:"num_reviews" bson:"num_reviews"`
	Sm2N             int                `json:"sm2_n" bson:"sm2_n"`
	Sm2EF            float64            `json:"sm2_ef" bson:"sm2_ef"`
	Sm2I             int                `json:"sm2_i" bson:"sm2_i"`
	CardType         int                `json:"card_type"`
}

func (card *Card) SetDefault() *Card {
	card.CreatedAt = time.Now()
	card.NextReview = card.CreatedAt
	card.NumReviews = 0
	card.Sm2N = 0
	card.Sm2EF = 2.5
	card.Sm2I = 0
	return card
}
