package dto

import "vietime-backend/internal/entity"

type GetAllDataResponse struct {
	AccessToken          string                       `json:"access_token"`
	RefreshToken         string                       `json:"refresh_token"`
	User                 entity.User                  `json:"user"`
	UserDeckAndCards     []entity.DeckWithCards       `json:"user_decks"`
	PublicDeckAndCards   []entity.DeckWithCards       `json:"public_decks"`
	DecksWithReviewCards []entity.DeckWithReviewCards `json:"decks_review"`
	AllCards             []entity.Card                `json:"all_cards"`
}
