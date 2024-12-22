package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// UpdateReviewCards	godoc
// UpdateReviewCards	API
//
//	@Summary		Update Review Cards
//	@Description	Update Review Cards
//	@Tags			card
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/card/review [put]
//	@Param			update_review_cards_request	body		dto.UpdateReviewCardsRequest	true	"Update Review Cards Request"
//	@Success		200							{object}	dto.UpdateReviewCardsResponse
//	@Failure		400							{object}	dto.ErrorResponse
//	@Failure		500							{object}	dto.ErrorResponse
func (h *restHandler) UpdateReviewCards(c *gin.Context) {
	var (
		req dto.UpdateReviewCardsRequest
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

	if len(req.CardIDs) == 0 || len(req.CardIDs) != len(req.IsCorrect) {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid card_ids[] or is_correct[] parameters"})
		return
	}

	deck.UpdateReview()
	cards, err := h.cardUseCase.GetCardsByDeck(&deckID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	needUpdate := make(map[string]bool)
	cardsMap := make(map[string]*entity.Card)
	for i := range *cards {
		cardsMap[(*cards)[i].ID.Hex()] = &(*cards)[i]
	}
	for i, id := range req.CardIDs {
		correct := req.IsCorrect[i]
		card1, exists := cardsMap[id.Hex()]
		if !exists {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Some card doesn't exist in given deck!"})
			return
		}
		card1.UpdateScheduleSM2(correct)
		needUpdate[id.Hex()] = true
		if correct {
			if card1.NumReviews == 1 {
				deck.CurNewCards++
				deck.TotalLearnedCards++
			} else {
				deck.CurReviewCards++
			}
		}
	}

	for id, update := range needUpdate {
		if update {
			err := h.cardUseCase.UpdateCardReview(cardsMap[id])
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
				return
			}
		}
	}

	cards, numBlueCards, numRedCards, numGreenCards, err := h.cardUseCase.GetReviewCardsByDeck(&deckID, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	updateDeckReq := &dto.UpdateDeckRequest{
		LastReview:        &deck.LastReview,
		CurNewCards:       &deck.CurNewCards,
		MaxNewCards:       &deck.MaxNewCards,
		TotalLearnedCards: &deck.TotalLearnedCards,
	}
	deck, err = h.deckUseCase.UpdateDeck(&deckID, updateDeckReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	_, err = h.userUseCase.AddXPToUser(&uID, req.TotalXP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUseCase.GetUserByID(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.UpdateReviewCardsResponse{
		User:          user,
		Cards:         *cards,
		NumBlueCards:  numBlueCards,
		NumRedCards:   numRedCards,
		NumGreenCards: numGreenCards,
	}
	c.JSON(http.StatusOK, resp)
}
