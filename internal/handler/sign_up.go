package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	config "vietime-backend/config"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// SignUp	godoc
// SignUp	API
//
//	@Summary		Sign Up
//	@Description	Sign Up
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/signup [post]
//	@Param			signup_request	formData	dto.SignupRequest	true	"Sign Up Request"
//	@Success		200				{object}	dto.SignupResponse
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		409				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
func (h *restHandler) SignUp(c *gin.Context) {
	var request dto.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.signUpUseCase.GetUserByEmail(&request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, dto.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user = &entity.User{
		Name:           request.Name,
		Email:          request.Email,
		HashedPassword: request.Password,
	}

	_, err = h.signUpUseCase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, er := h.signUpUseCase.CreateAccessToken(user, &config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if er != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.signUpUseCase.CreateRefreshToken(user, &config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := dto.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
