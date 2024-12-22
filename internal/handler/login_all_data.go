package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	bootstrap "vietime-backend/config"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// LogInGetAllData	godoc
// LogInGetAllData	API
//
//	@Summary		Log In And Get All Data
//	@Description	Log In And Get All Data
//	@Tags			mobile
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/login-get-all [post]
//	@Param			login_request	formData	dto.LoginRequest	true	"Log In Request"
//	@Success		200				{object}	dto.LoginGetAllDataResponse
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		401				{object}	dto.ErrorResponse
//	@Failure		404				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
func (h *restHandler) LogInGetAllData(c *gin.Context) {
	var request dto.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.signUpUseCase.GetUserByEmail(&request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Invalid credentials"})
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

	uID := user.ID.Hex()
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

	loginResponse := dto.LoginGetAllDataResponse{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		User:                 *user,
		UserDeckAndCards:     *userDecks,
		PublicDeckAndCards:   *publicDecks,
		DecksWithReviewCards: *decksWithReviewCard,
		AllCards:             allCards,
	}

	c.JSON(http.StatusOK, loginResponse)
}
