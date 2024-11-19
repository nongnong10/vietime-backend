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
