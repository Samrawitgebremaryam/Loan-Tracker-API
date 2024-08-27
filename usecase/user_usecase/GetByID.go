package user_usecase

import (
	"context"
	"loan_tracker_api/domain"
)

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetUserByID(ctx, id)
}
