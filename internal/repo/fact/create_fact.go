package fact

import (
	"context"
	"vietime-backend/internal/entity"
)

func (fr *factRepo) CreateFact(fact *entity.Fact) error {
	_, err := fr.db.Collection("fun_facts").InsertOne(context.TODO(), *fact)
	if err != nil {
		return err
	}

	return nil
}
