package usecase

import (
	"vietime-backend/internal/entity"
	user "vietime-backend/internal/repo/user"
)

type SignUpUseCase interface {
	Create(user *entity.User) (string, error)
	GetUserByEmail(email *string) (*entity.User, error)
	CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error)
}

type signUpUseCase struct {
	userRepository user.UserRepository
}

func NewSignUpUseCase(userRepository user.UserRepository) SignUpUseCase {
	return &signUpUseCase{
		userRepository: userRepository,
	}
}
