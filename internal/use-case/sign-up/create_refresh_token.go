package usecase

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/utils"
)

func (su *signupUsecase) CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
