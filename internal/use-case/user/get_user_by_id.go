package user

import "vietime-backend/internal/entity"

func (uc *userUseCase) GetUserByID(id *string) (*entity.User, error) {
	return uc.userRepository.GetByID(id)
}
