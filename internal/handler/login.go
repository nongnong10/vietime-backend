package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	config "vietime-backend/config"
	"vietime-backend/internal/delivery/http/dto"
)

// Login    godoc
// Login	API
//
//	@Summary		Login
//	@Description	Login
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/login [post]
//	@Param			login_request	formData	dto.LoginRequest	true	"Login Request"
//	@Success		200				{object}	dto.LoginResponse
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		409				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
func (h *restHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	// 1. Check if user is existed with the given email
	user, err := h.signUpUseCase.GetUserByEmail(&loginRequest.Email)
	if user == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "User is not found with the given email",
		})
		return
	}

	// 2. Check password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "Password is incorrect",
		})
		return
	}

	// 3. Generate JWT token
	accessToken, err := h.signUpUseCase.CreateAccessToken(user, &config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Failed to generate access token",
		})
	}
	refreshToken, err := h.signUpUseCase.CreateRefreshToken(user, &config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Failed to generate refresh token",
		})
	}

	loginResponse := dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, loginResponse)
}
