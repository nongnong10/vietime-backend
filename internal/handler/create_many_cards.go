package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/utils/slice"
)

// CreateManyCards	godoc
// CreateManyCards	API
//
//	@Summary		Create Many Cards
//	@Description	Create Many Cards
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/create_list [post]
//	@Param			create_card_request	body		dto.CreateManyCardsRequest	true	"Create List New Cards"
//	@Success		200					{object}	dto.CreateManyCardsResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) CreateManyCards(ctx *gin.Context) {
	var createCardsRequest dto.CreateManyCardsRequest

	// 1. Get userID from context
	userID, err := GetLoggedInUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: fmt.Sprintf("Cannot get userID from context: %s", err.Error()),
		})
		return
	}

	// 2. Bind request
	createCardsRequest.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}
	err = ctx.ShouldBind(&createCardsRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// 3. Validate request
	for _, card := range createCardsRequest.Cards {
		if len(card.WrongAnswers) < 3 {
			ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: fmt.Sprintf("WrongAnswers have %v answers, WrongAnswers must have at least 3 items", len(card.WrongAnswers))})
			return
		}
	}

	// 4. Create cards
	cards := slice.Map(createCardsRequest.Cards, func(card dto.Card) entity.Card {
		return entity.Card{
			UserID:           createCardsRequest.UserID,
			DeckID:           createCardsRequest.DeckID,
			QuestionImgURL:   card.QuestionImgURL,
			QuestionImgLabel: card.QuestionImgLabel,
			Question:         card.Question,
			Answer:           card.Answer,
			WrongAnswers:     card.WrongAnswers,
		}
	})

	err = h.cardUseCase.CreateManyCards(cards)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// 5. Response
	ctx.JSON(http.StatusOK, dto.CreateManyCardsResponse{
		Success: true,
	})
}
