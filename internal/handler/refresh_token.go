package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	bootstrap "vietime-backend/config"
	"vietime-backend/internal/delivery/http/dto"
)

// RefreshToken	godoc
// RefreshToken	API
//
//	@Summary		Refresh Token
//	@Description	Refresh Token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/api/refresh [post]
//	@Param			refresh_token_request	body		dto.RefreshTokenRequest	true	"Refresh Token Request"
//	@Success		200						{object}	dto.RefreshTokenResponse
//	@Failure		401						{object}	dto.ErrorResponse
//	@Failure		500						{object}	dto.ErrorResponse
func (h *restHandler) RefreshToken(c *gin.Context) {
	var request dto.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := h.signUpUseCase.ExtractIDFromToken(&request.RefreshToken, &bootstrap.E.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUseCase.GetUserByID(&id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
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

	refreshTokenResponse := dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
