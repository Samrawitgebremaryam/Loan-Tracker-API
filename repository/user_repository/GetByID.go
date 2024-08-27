package user_repository

import (
	"context"
	"loan_tracker_api/domain"
	"loan_tracker_api/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *userRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	objectID, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	err = r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
