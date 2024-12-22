package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// GetFact	godoc
// GetFact	API
//
//	@Summary		Get Fact
//	@Description	Get Fact
//	@Tags			fact
//	@Produce		json
//	@Router			/api/fact [get]
//	@Success		200	{object}	dto.GetFactResponse
//	@Failure		500	{object}	dto.ErrorResponse
func (h *restHandler) GetFact(c *gin.Context) {
	fact, err := h.factUseCase.GetFact()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}
	resp := dto.GetFactResponse{
		Fact: fact.Content,
	}
	c.JSON(http.StatusOK, resp)
}
