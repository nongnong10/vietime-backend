package usecase

import "vietime-backend/pkg/utils/token"

func (suu *signUpUseCase) ExtractIDFromToken(requestToken *string, secret *string) (string, error) {
	return token.ExtractIDFromToken(requestToken, secret)
}
