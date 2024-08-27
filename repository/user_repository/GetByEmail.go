package user_repository

import (
	"context"
	"loan_tracker_api/domain"
	"loan_tracker_api/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
