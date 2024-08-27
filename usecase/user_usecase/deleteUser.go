package user_usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *userUsecase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.DeleteUser(ctx, id)
}
