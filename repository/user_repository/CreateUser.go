package user_repository

import (
	"context"
	"loan_tracker_api/domain"
)

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
