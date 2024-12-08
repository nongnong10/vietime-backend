package user

import (
	"vietime-backend/internal/entity"
	user "vietime-backend/internal/repo/user"
)

type UserUseCase interface {
	GetUserByID(id *string) (*entity.User, error)
}

type userUseCase struct {
	userRepository user.UserRepository
}

func NewUserUseCase(userRepository user.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
