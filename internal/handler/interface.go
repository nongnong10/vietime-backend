package handler

import (
	"errors"
	signUpUseCase "vietime-backend/internal/use-case/sign-up"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type restHandler struct {
	signUpUseCase signUpUseCase.SignUpUseCase
}

func NewHandler(signUpUseCase signUpUseCase.SignUpUseCase) RestHandler {
	return &restHandler{
		signUpUseCase: signUpUseCase,
	}
}

func GetLoggedInUserID(c *gin.Context) (string, error) {
	uID, isExisted := c.Get("x-user-id")
	if !isExisted {
		return "", errors.New("Missing x-user-id (set at middleware) in Gin Context")
	}

	return uID.(string), nil
}
