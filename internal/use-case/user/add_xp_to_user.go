package user

import "vietime-backend/internal/entity"

func (uu *userUseCase) AddXPToUser(userID *string, XP int) (*entity.User, error) {
	user, err := uu.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	user.XP += XP
	user.UpdateLevel()
	user.UpdateStreak()
	err = uu.userRepository.UpdateUserXP(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
