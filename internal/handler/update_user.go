package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"vietime-backend/internal/delivery/http/dto"
)

// UpdateUser	godoc
// UpdateUser	API
//
//	@Summary		Update User Details
//	@Description	Update User Details
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Router			/api/user/update [put]
//	@Param			update_user_request	body		dto.UpdateUserRequest	true	"Update User Request"
//	@Success		200					{object}	dto.UpdateUserResponse
//	@Failure		400					{object}	dto.ErrorResponse
//	@Failure		500					{object}	dto.ErrorResponse
func (h *restHandler) UpdateUser(c *gin.Context) {
	var (
		req dto.UpdateUserRequest
		err error
	)

	uID, err := GetLoggedInUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	if req.OldPassword != nil && req.NewPassword != nil {
		user, err := h.userUseCase.GetUserByID(&uID)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(*req.OldPassword)) != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Wrong old password! Couldn't update password!"})
			return
		}
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(*req.NewPassword),
			bcrypt.DefaultCost,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
			return
		}
		pass := string(encryptedPassword)
		req.HashedPassword = &pass
		req.OldPassword = nil
		req.NewPassword = nil
	}

	user, err := h.userUseCase.UpdateUser(&uID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	resp := dto.UpdateUserResponse{
		User: *user,
	}
	c.JSON(http.StatusOK, resp)
}
