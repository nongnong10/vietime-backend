package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

// CreateFact	godoc
// CreateFact	API
//
//	@Summary		Create New Fact
//	@Description	Create New Fact
//	@Tags			fact
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/fact/create [post]
//	@Param			create_fact_request	body		dto.CreateFactRequest	true	"Create Fact Request"
//	@Success		200					{object}	dto.CreateFactResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) CreateFact(c *gin.Context) {
	var (
		req dto.CreateFactRequest
		err error
	)

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	fact := &entity.Fact{
		Content: req.Content,
	}
	err = h.factUseCase.CreateFact(fact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	createFactResponse := dto.CreateFactResponse{
		Success: true,
	}

	c.JSON(http.StatusOK, createFactResponse)
}
