package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	config "vietime-backend/config"
	"vietime-backend/internal/delivery/http/request"
	"vietime-backend/internal/delivery/http/response"
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
//	@Param			login_request	formData	request.LoginRequest	true	"Login Request"
//	@Success		200				{object}	response.LoginResponse
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		409				{object}	response.ErrorResponse
//	@Failure		500				{object}	response.ErrorResponse
func (h *restHandler) Login(c *gin.Context) {
	var loginRequest request.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	// 1. Check if user is existed with the given email
	user, err := h.signUpUsecase.GetUserByEmail(&loginRequest.Email)
	if user == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: "User is not found with the given email",
		})
		return
	}

	// 2. Check password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: "Password is incorrect",
		})
		return
	}

	// 3. Generate JWT token
	accessToken, err := h.signUpUsecase.CreateAccessToken(user, &config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Failed to generate access token",
		})
	}
	refreshToken, err := h.signUpUsecase.CreateRefreshToken(user, &config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Failed to generate refresh token",
		})
	}

	loginResponse := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, loginResponse)
}
