package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"vietime-backend/internal/entity"
)

func (ur *userRepository) UpdateUserXP(user *entity.User) error {
	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "xp", Value: user.XP},
			{Key: "xp_to_level_up", Value: user.XPToLevelUp},
			{Key: "level", Value: user.Level},
			{Key: "streak", Value: user.Streak},
			{Key: "last_streak", Value: user.LastStreak},
		}},
	}
	_, err := ur.db.Collection(ur.colName).UpdateOne(context.TODO(), &filter, &update)
	if err != nil {
		return err
	}
	return nil
}
