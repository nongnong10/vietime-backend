package user

import (
	"vietime-backend/internal/delivery/http/dto"
	"vietime-backend/internal/entity"
)

func (uc *userUseCase) UpdateUser(userID *string, req *dto.UpdateUserRequest) (*entity.User, error) {
	return uc.userRepository.UpdateUser(userID, req)
}
