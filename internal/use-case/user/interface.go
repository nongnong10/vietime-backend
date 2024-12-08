package user

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
	user "vietime-backend/internal/repo/user"
)

type UserUseCase interface {
	GetUserByID(id *string) (*entity.User, error)
	UpdateUser(userID *string, req *dto.UpdateUserRequest) (*entity.User, error)
}

type userUseCase struct {
	userRepository user.UserRepository
}

func NewUserUseCase(userRepository user.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
