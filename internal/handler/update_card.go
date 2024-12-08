package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// UpdateCard	godoc
// UpdateCard	API
//
//	@Summary		Update Card Details
//	@Description	Update Card Details
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/update [put]
//	@Param			update_card_request	body		dto.UpdateCardRequest	true	"Update Card Request"
//	@Success		200					{object}	dto.UpdateCardResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) UpdateCard(c *gin.Context) {
	var (
		req dto.UpdateCardRequest
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
	if card.UserID.Hex() != uID {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Not your card! Can't update! Logged in user != card's user"})
		return
	}

	req.CardID = nil
	card, err = h.cardUseCase.UpdateCard(&cardID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.UpdateCardResponse{
		Success: true,
		Card:    *card,
	}
	c.JSON(http.StatusOK, resp)
}
