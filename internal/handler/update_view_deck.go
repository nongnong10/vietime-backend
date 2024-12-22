package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// UpdateViewDeck	godoc
// UpdateViewDeck	API
//
//	@Summary		Update View Deck
//	@Description	Update View Deck
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Router			/api/deck/view [put]
//	@Param			view_deck_request	body		dto.UpdateViewDeckRequest	true	"View Deck Request"
//	@Success		200					{object}	dto.SuccessResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) UpdateViewDeck(c *gin.Context) {
	var (
		updateViewDeckReq dto.UpdateViewDeckRequest
		updateDeckReq     dto.UpdateDeckRequest
		err               error
	)

	err = c.ShouldBind(&updateViewDeckReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	deckID := updateViewDeckReq.DeckID.Hex()

	updateDeckReq = dto.UpdateDeckRequest{
		Views: updateViewDeckReq.Views,
	}
	_, err = h.deckUseCase.UpdateDeck(&deckID, &updateDeckReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.SuccessResponse{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}
