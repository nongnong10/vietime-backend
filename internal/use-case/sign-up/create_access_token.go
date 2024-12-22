package usecase

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/utils/token"
)

func (su *signUpUseCase) CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error) {
	return token.CreateAccessToken(user, secret, expiry)
}
