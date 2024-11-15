package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vietime-backend/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) (string, error)
	GetByEmail(email *string) (*entity.User, error)
}

type userRepository struct {
	db      *mongo.Database
	colName string
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db:      db,
		colName: "users",
	}
}
