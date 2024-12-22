package deck

func (uc *deckUsecase) DeleteDeck(deckID *string) error {
	return uc.deckRepo.DeleteDeck(deckID)
}
