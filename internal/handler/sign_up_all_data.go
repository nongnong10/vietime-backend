package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	bootstrap "vietime-backend/config"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// SignUpGetAllData	godoc
// SignUpGetAllData	API
//
//	@Summary		Sign Up And Get All Data
//	@Description	Sign Up And Get All Data
//	@Tags			mobile
//	@Accept			multipart/form-data
//	@Produce		json
//	@Router			/api/signup-get-all [post]
//	@Param			signup_request	formData	dto.SignupRequest	true	"Sign Up Request"
//	@Success		200				{object}	dto.SignUpGetAllDataResponse
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		401				{object}	dto.ErrorResponse
//	@Failure		404				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
func (h *restHandler) SignUpGetAllData(c *gin.Context) {
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

	uID, err := h.signUpUseCase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	user.ID, _ = primitive.ObjectIDFromHex(uID)

	accessToken, er := h.signUpUseCase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if er != nil {
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

	signupResponse := dto.SignUpGetAllDataResponse{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		User:                 *user,
		UserDeckAndCards:     *userDecks,
		PublicDeckAndCards:   *publicDecks,
		DecksWithReviewCards: *decksWithReviewCard,
		AllCards:             allCards,
	}

	c.JSON(http.StatusOK, signupResponse)
}
