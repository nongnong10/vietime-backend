package usecase

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/utils"
)

func (su *signUpUseCase) CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}
