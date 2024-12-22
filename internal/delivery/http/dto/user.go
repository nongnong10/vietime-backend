package dto

import "vietime-backend/internal/entity"

// Login

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginGetAllDataResponse struct {
	AccessToken          string                       `json:"access_token"`
	RefreshToken         string                       `json:"refresh_token"`
	User                 entity.User                  `json:"user"`
	UserDeckAndCards     []entity.DeckWithCards       `json:"user_decks"`
	PublicDeckAndCards   []entity.DeckWithCards       `json:"public_decks"`
	DecksWithReviewCards []entity.DeckWithReviewCards `json:"decks_review"`
	AllCards             []entity.Card                `json:"all_cards"`
}

// Sign Up

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Refresh Token

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Update User

type UpdateUserRequest struct {
	Name           *string `json:"name" bson:"name,omitempty"`
	OldPassword    *string `json:"old_password" bson:"old_password,omitempty"`
	NewPassword    *string `json:"new_password" bson:"new_password,omitempty"`
	HashedPassword *string `json:"hashed_password" bson:"hashed_password,omitempty" swaggerignore:"true"`
}

type UpdateUserResponse struct {
	User entity.User `json:"user"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}
