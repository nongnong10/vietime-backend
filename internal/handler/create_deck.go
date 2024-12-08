package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// CreateDeck	godoc
// CreateDeck	API
//
//	@Summary		Create New Deck
//	@Description	Create New Deck
//	@Tags			deck
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/deck/create [post]
//	@Param			create_deck_request	body		dto.CreateDeckRequest	true	"Create Deck Request"
//	@Success		200					{object}	dto.CreateDeckResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) CreateDeck(c *gin.Context) {
	var (
		req dto.CreateDeckRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	req.UserID, err = primitive.ObjectIDFromHex(uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.userUseCase.GetUserByID(&uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "User ID doesn't exist in DB"})
		return
	}
	deck := &entity.Deck{
		UserID:              req.UserID,
		Name:                req.Name,
		Description:         req.Description,
		DescriptionImageURL: req.DescriptionImageURL,
		Position:            req.Position,
		TotalCards:          req.TotalCards,
	}
	deck, err = h.deckUseCase.CreateDeck(deck)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	createDeckResponse := dto.CreateDeckResponse{
		Deck: *deck,
	}

	c.JSON(http.StatusOK, createDeckResponse)
}
