package user_usecase

import (
	"loan_tracker_api/domain"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	authService    domain.AuthService
	emailService   domain.EmailService
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, authService domain.AuthService, emailService domain.EmailService, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       userRepository,
		emailService:   emailService,
		authService:    authService,
		contextTimeout: timeout,
	}
}
