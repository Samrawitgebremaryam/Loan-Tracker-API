package user_usecase

import (
	"context"
	"loan_tracker_api/domain"
)

func (u *userUsecase) GetUsers(ctx context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetUsers(ctx)
}
