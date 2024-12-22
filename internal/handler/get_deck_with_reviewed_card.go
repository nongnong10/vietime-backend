package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// GetDeckWithReviewCards	godoc
// GetDeckWithReviewCards	API
//
//	@Summary		Get Deck With Review Cards Of Logged In User
//	@Description	Get Deck With Review Cards Of Logged In User
//	@Tags			deck
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/review-cards [get]
//	@Success		200	{object}	[]entity.DeckWithReviewCards
//	@Failure		500	{object}	dto.ErrorResponse
func (h *restHandler) GetDeckWithReviewCards(c *gin.Context) {
	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	deckWithCards, err := h.deckUseCase.GetReviewCardsAllDecksOfUser(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, deckWithCards)
}
