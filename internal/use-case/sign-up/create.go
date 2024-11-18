package usecase

import "vietime-backend/internal/entity"

func (su *signUpUseCase) Create(user *entity.User) (string, error) {
	return su.userRepository.Create(user)
}
