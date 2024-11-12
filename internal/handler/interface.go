package handler

import (
	"errors"
	usecase "vietime-backend/internal/use-case/sign-up"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}

type restHandler struct {
	signUpUsecase usecase.SignupUsecase
}

func NewHandler(signUpUseCase usecase.SignupUsecase) RestHandler {
	return &restHandler{
		signUpUsecase: signUpUseCase,
	}
}

func GetLoggedInUserID(c *gin.Context) (string, error) {
	uID, isExisted := c.Get("x-user-id")
	if !isExisted {
		return "", errors.New("Missing x-user-id (set at middleware) in Gin Context")
	}

	return uID.(string), nil
}
