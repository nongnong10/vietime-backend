package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	bootstrap "vietime-backend/config"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// GetAllData	godoc
// GetAllData	API
//
//	@Summary		Get All Data
//	@Description	Get All Data
//	@Tags			mobile
//	@Accept			json
//	@Produce		json
//	@Router			/api/get-all [post]
//	@Param			refresh_token_request	body		dto.RefreshTokenRequest	true	"Refresh Token Request"
//	@Success		200						{object}	dto.GetAllDataResponse
//	@Failure		500						{object}	dto.ErrorResponse
func (h *restHandler) GetAllData(c *gin.Context) {
	var request dto.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	uID, err := h.signUpUseCase.ExtractIDFromToken(&request.RefreshToken, &bootstrap.E.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUseCase.GetUserByID(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := h.signUpUseCase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.signUpUseCase.CreateRefreshToken(user, &bootstrap.E.RefreshTokenSecret, bootstrap.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	userDecks, publicDecks, decksWithReviewCard, err := h.deckUseCase.GetDecksWithCards(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	allCards := []entity.Card{}
	for i := range *userDecks {
		allCards = append(allCards, *(*userDecks)[i].Cards...)
	}
	for i := range *publicDecks {
		allCards = append(allCards, *(*publicDecks)[i].Cards...)
	}

	getAllResponse := dto.GetAllDataResponse{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		User:                 *user,
		UserDeckAndCards:     *userDecks,
		PublicDeckAndCards:   *publicDecks,
		DecksWithReviewCards: *decksWithReviewCard,
		AllCards:             allCards,
	}

	c.JSON(http.StatusOK, getAllResponse)
}
