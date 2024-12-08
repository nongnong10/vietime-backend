package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// UpdateDeck	godoc
// UpdateDeck	API
//
//	@Summary		Update Deck Details
//	@Description	Update Deck Details
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/update [put]
//	@Param			update_deck_request	body		dto.UpdateDeckRequest	true	"Update Deck Request"
//	@Success		200					{object}	dto.UpdateDeckResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) UpdateDeck(c *gin.Context) {
	var (
		req dto.UpdateDeckRequest
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

	req.DeckID = nil
	deck, err = h.deckUseCase.UpdateDeck(&deckID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.UpdateDeckResponse{
		Success: true,
		Deck:    *deck,
	}
	c.JSON(http.StatusOK, resp)
}
