package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// DeleteCard	godoc
// DeleteCard	API
//
//	@Summary		Delete Card
//	@Description	Delete Card
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/delete [delete]
//	@Param			delete_card_request	body		dto.DeleteCardRequest	true	"Delete Card Request"
//	@Success		200					{object}	dto.DeleteCardResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) DeleteCard(c *gin.Context) {
	var (
		req dto.DeleteCardRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	cardID := req.CardID.Hex()
	card, err := h.cardUseCase.GetCardByID(&cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	deckID := card.DeckID.Hex()
	deck, err := h.deckUseCase.GetDeckByID(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	if deck.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Not your card! Can't update! Logged in user != card's user"})
		return
	}

	err = h.cardUseCase.DeleteCard(&cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	deleteCardResponse := dto.DeleteCardResponse{
		Success: true,
	}

	c.JSON(http.StatusOK, deleteCardResponse)
}
