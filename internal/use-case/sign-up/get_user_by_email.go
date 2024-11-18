package usecase

import "vietime-backend/internal/entity"

func (su *signUpUseCase) GetUserByEmail(email *string) (*entity.User, error) {
	return su.userRepository.GetByEmail(email)
}
