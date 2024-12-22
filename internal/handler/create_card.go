package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// CreateCard	godoc
// CreateCard	API
//
//	@Summary		Create New Card
//	@Description	Create New Card
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/create [post]
//	@Param			create_card_request	body		dto.CreateCardRequest	true	"Create Card Request"
//	@Success		200					{object}	dto.CreateCardResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) CreateCard(ctx *gin.Context) {
	var createCardRequest dto.CreateCardRequest

	// 1. Get userID from context
	userID, err := GetLoggedInUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: fmt.Sprintf("Cannot get userID from context: %s", err.Error()),
		})
		return
	}

	// 2. Bind request
	createCardRequest.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}
	err = ctx.ShouldBind(&createCardRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// 3. Validate request
	if len(createCardRequest.WrongAnswers) < 3 {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: fmt.Sprintf("WrongAnswers have %v answers, WrongAnswers must have at least 3 items", len(createCardRequest.WrongAnswers))})
		return
	}

	// 4. Create card
	card := &entity.Card{
		UserID:           createCardRequest.UserID,
		DeckID:           createCardRequest.DeckID,
		QuestionImgURL:   createCardRequest.QuestionImgURL,
		QuestionImgLabel: createCardRequest.QuestionImgLabel,
		Question:         createCardRequest.Question,
		Answer:           createCardRequest.Answer,
		WrongAnswers:     createCardRequest.WrongAnswers,
	}
	card, err = h.cardUseCase.CreateCard(card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// 5. Response
	ctx.JSON(http.StatusOK, dto.CreateCardResponse{
		Card: *card,
	})
}
