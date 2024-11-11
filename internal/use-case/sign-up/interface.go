package usecase

import (
	"vietime-backend/internal/entity"
	repository "vietime-backend/internal/repo/user"
)

type SignupUsecase interface {
	Create(user *entity.User) (string, error)
	GetUserByEmail(email *string) (*entity.User, error)
	CreateAccessToken(user *entity.User, secret *string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret *string, expiry int) (refreshToken string, err error)
}

type signupUsecase struct {
	userRepository repository.UserRepository
}

func NewSignupUsecase(userRepository repository.UserRepository) SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}
