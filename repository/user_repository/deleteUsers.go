package user_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *userRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	objectID, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return err
	}

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if res == 0 {
		return errors.New("user not found")
	}
	return nil
}
