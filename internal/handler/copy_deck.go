package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// CopyDeck	godoc
// CopyDeck	API
//
//	@Summary		Copy Deck
//	@Description	Copy Deck
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/copy [post]
//	@Param			copy_deck_request	body		dto.CopyDeckRequest	true	"Copy Deck Request"
//	@Success		200					{object}	dto.CopyDeckResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		401					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) CopyDeck(c *gin.Context) {
	var (
		req dto.CopyDeckRequest
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
	if !deck.IsPublic && deck.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Can't copy private deck! Deck is not public and Logged in user != deck's user"})
		return
	}

	deckWithCards, deckWithReviewCards, err := h.deckUseCase.CopyDeck(&uID, &deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.CopyDeckResponse{
		Deck:       *deckWithCards,
		DeckReview: *deckWithReviewCards,
	}
	c.JSON(http.StatusOK, resp)
}
