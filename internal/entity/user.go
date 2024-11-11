package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	LEVEL_XP_INC = 100
)

type User struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	Name           string             `json:"name" bson:"name"`
	Email          string             `json:"email" bson:"email"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	XP             int                `json:"xp" bson:"xp"`
	XPToLevelUp    int                `json:"xp_to_level_up" bson:"xp_to_level_up"`
	Level          int                `json:"level" bson:"level"`
	Streak         int                `json:"streak" bson:"streak"`
	LastStreak     time.Time          `json:"last_streak" bson:"last_streak"`
	IsAdmin        bool               `json:"is_admin" bson:"is_admin"`
}

func (user *User) SetDefault() *User {
	user.CreatedAt = time.Now()
	user.XP = 0
	user.XPToLevelUp = 100
	user.Level = 1
	user.Streak = 1
	user.LastStreak = user.CreatedAt
	user.IsAdmin = false
	return user
}
