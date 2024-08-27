package user_usecase

import (
	"context"
	"loan_tracker_api/domain"
)

func (uc *userUsecase) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	u, err := uc.userRepo.GetByUsernameOrEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return *u, nil
}
