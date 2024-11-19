package handler

import (
	"errors"
	"fmt"
	"vietime-backend/internal/use-case/card"
	sign_up "vietime-backend/internal/use-case/sign-up"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	CreateCard(ctx *gin.Context)
}

type restHandler struct {
	signUpUseCase sign_up.SignUpUseCase
	cardUseCase   card.CardUseCase
}

func NewHandler(signUpUseCase sign_up.SignUpUseCase) RestHandler {
	return &restHandler{
		signUpUseCase: signUpUseCase,
	}
}

func GetLoggedInUserID(c *gin.Context) (string, error) {
	uID, isExisted := c.Get("x-user-id")
	if !isExisted {
		return "", errors.New(fmt.Sprintf("Missing x-user-id (set at middleware) in Gin Context: %v", c.Request.Header))
	}

	return uID.(string), nil
}
