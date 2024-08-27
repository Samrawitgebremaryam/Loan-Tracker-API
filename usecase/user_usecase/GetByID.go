package user_usecase

import (
	"context"
	"loan_tracker_api/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *userUsecase) GetUserByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetUserByID(ctx, id)
}
