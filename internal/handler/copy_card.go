package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// CopyCardToDeck	godoc
// CopyCardToDeck	API
//
//	@Summary		Copy Card To Deck
//	@Description	Copy Card To Deck
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/copy [post]
//	@Param			copy_card_to_deck_request	body		dto.CopyCardToDeckRequest	true	"Copy Card To Deck Request"
//	@Success		200							{object}	dto.CopyCardToDeckResponse
//	@Failure		400							{object}	dto.ErrorResponse
//	@Failure		500							{object}	dto.ErrorResponse
func (h *restHandler) CopyCardToDeck(c *gin.Context) {
	var (
		req dto.CopyCardToDeckRequest
		err error
	)

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	cardID := req.CardID.Hex()
	deckID := req.DeckID.Hex()
	card, err := h.cardUseCase.CopyCardToDeck(&cardID, &deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.CopyCardToDeckResponse{
		Card: *card,
	}
	c.JSON(http.StatusOK, resp)
}
