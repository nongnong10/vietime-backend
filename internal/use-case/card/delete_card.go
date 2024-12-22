package card

func (uc *cardUseCase) DeleteCard(cardID *string) error {
	return uc.cardRepo.DeleteCard(cardID)
}
