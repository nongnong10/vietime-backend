package handler

import (
	"errors"
	"fmt"
	"vietime-backend/internal/use-case/card"
	"vietime-backend/internal/use-case/deck"
	sign_up "vietime-backend/internal/use-case/sign-up"
	"vietime-backend/internal/use-case/user"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	UpdateUser(c *gin.Context)

	CreateCard(ctx *gin.Context)
	UpdateCard(c *gin.Context)

	CreateDeck(c *gin.Context)
	UpdateDeck(c *gin.Context)
}

type restHandler struct {
	signUpUseCase sign_up.SignUpUseCase
	cardUseCase   card.CardUseCase
	userUseCase   user.UserUseCase
	deckUseCase   deck.DeckUsecase
}

func NewHandler(
	signUpUseCase sign_up.SignUpUseCase,
	cardUseCase card.CardUseCase,
	userUseCase user.UserUseCase,
	deckUseCase deck.DeckUsecase,
) RestHandler {
	return &restHandler{
		signUpUseCase: signUpUseCase,
		cardUseCase:   cardUseCase,
		userUseCase:   userUseCase,
		deckUseCase:   deckUseCase,
	}
}

func GetLoggedInUserID(c *gin.Context) (string, error) {
	uID, isExisted := c.Get("x-user-id")
	if !isExisted {
		return "", errors.New(fmt.Sprintf("Missing x-user-id (set at middleware) in Gin Context: %v", c.Request.Header))
	}

	return uID.(string), nil
}
