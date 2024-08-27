package user_usecase

import "context"

func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.DeleteUser(ctx, id)
}
