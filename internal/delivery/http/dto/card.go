package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vietime-backend/internal/entity"
)

type CreateCardRequest struct {
	UserID           primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	DeckID           primitive.ObjectID `json:"deck_id" binding:"required"`
	Index            int                `json:"index"`
	QuestionImgURL   string             `json:"question_img_url"`
	QuestionImgLabel string             `json:"question_img_label"`
	Question         string             `json:"question" binding:"required"`
	Answer           string             `json:"answer" binding:"required"`
	WrongAnswers     []string           `json:"wrong_answers" binding:"required"`
}

type CreateCardResponse struct {
	Card entity.Card `json:"card"`
}

// Update Card
type UpdateCardRequest struct {
	CardID           *primitive.ObjectID `json:"card_id" bson:"_id,omitempty" binding:"required"`
	DeckID           *primitive.ObjectID `json:"deck_id" bson:"deck_id,omitempty"`
	QuestionImgURL   *string             `json:"question_img_url" bson:"question_img_url,omitempty"`
	QuestionImgLabel string              `json:"question_img_label" bson:"question_img_label,omitempty"`
	Question         *string             `json:"question" bson:"question,omitempty"`
	Answer           *string             `json:"answer" bson:"answer,omitempty"`
	WrongAnswers     *[]string           `json:"wrong_answers" bson:"wrong_answers,omitempty"`
}

type UpdateCardResponse struct {
	Success bool        `json:"success"`
	Card    entity.Card `json:"card"`
}
