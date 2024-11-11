package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	bootstrap "vietime-backend/config"
	"vietime-backend/internal/delivery/http/request"
	"vietime-backend/internal/delivery/http/response"
	"vietime-backend/internal/entity"
	usecase "vietime-backend/internal/use-case/sign-up"
)

type restHandler struct {
	signUpUsecase usecase.SignupUsecase
}

func NewHandler(signUpUseCase usecase.SignupUsecase) RestHandler {
	return &restHandler{
		signUpUsecase: signUpUseCase,
	}
}

func (h *restHandler) SignUp(c *gin.Context) {
	var request request.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.signUpUsecase.GetUserByEmail(&request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, response.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user = &entity.User{
		Name:           request.Name,
		Email:          request.Email,
		HashedPassword: request.Password,
	}

	_, err = h.signUpUsecase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, er := h.signUpUsecase.CreateAccessToken(user, &bootstrap.E.AccessTokenSecret, bootstrap.E.AccessTokenExpiryHour)
	if er != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.signUpUsecase.CreateRefreshToken(user, &bootstrap.E.RefreshTokenSecret, bootstrap.E.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := response.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
