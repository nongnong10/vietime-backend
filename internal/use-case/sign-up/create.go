package usecase

import "vietime-backend/internal/entity"

func (su *signupUsecase) Create(user *entity.User) (string, error) {
	return su.userRepository.Create(user)
}
