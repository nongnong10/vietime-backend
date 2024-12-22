package usecase

import (
	"vietime-backend/internal/entity"
	"vietime-backend/pkg/utils/token"
)

func (su *signUpUseCase) CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error) {
	return token.CreateRefreshToken(user, secret, expiry)
}
