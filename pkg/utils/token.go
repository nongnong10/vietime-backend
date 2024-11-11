package utils

import (
	"fmt"
	"time"
	"vietime-backend/internal/entity"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name string `json:"user_name"`
	ID   string `json:"user_id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *entity.User, secret *string, expireHours int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expireHours))
	claims := &JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(*secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *entity.User, secret *string, expireHours int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expireHours))
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(*secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized(requestToken *string, secret *string) (bool, error) {
	_, err := jwt.Parse(*requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken *string, secret *string) (string, error) {
	token, err := jwt.Parse(*requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["user_id"].(string), nil
}
