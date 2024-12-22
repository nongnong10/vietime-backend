package fact

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"vietime-backend/internal/entity"
)

func (fr *factRepo) GetFact() (*entity.Fact, error) {
	pipeline := bson.A{
		bson.D{{"$sample", bson.D{{"size", 1}}}},
	}

	opts := options.Aggregate().SetMaxTime(2 * time.Second) // You can adjust the timeout as needed

	cursor, err := fr.db.Collection("fun_facts").Aggregate(context.TODO(), pipeline, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if !cursor.Next(context.TODO()) {
		return nil, errors.New("no facts found")
	}

	var result entity.Fact
	err = cursor.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
