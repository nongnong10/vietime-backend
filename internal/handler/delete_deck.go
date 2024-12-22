package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// DeleteDeck	godoc
// DeleteDeck	API
//
//	@Summary		Delete Deck
//	@Description	Delete Deck
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/delete [delete]
//	@Param			delete_deck_request	body		dto.DeleteDeckRequest	true	"Delete Deck Request"
//	@Success		200					{object}	dto.DeleteDeckResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) DeleteDeck(c *gin.Context) {
	var (
		req dto.DeleteDeckRequest
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

	deckID := req.DeckID.Hex()
	deck, err := h.deckUseCase.GetDeckByID(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	if deck.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Not your deck! Can't update! Logged in user != deck's user"})
		return
	}

	err = h.deckUseCase.DeleteDeck(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	deleteDeckResponse := dto.DeleteDeckResponse{
		Success: true,
	}

	c.JSON(http.StatusOK, deleteDeckResponse)
}
