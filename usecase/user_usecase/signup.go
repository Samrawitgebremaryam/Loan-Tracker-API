package user_usecase

import (
	"context"
	"errors"
	"fmt"
	"loan_tracker_api/domain"
	"loan_tracker_api/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) SignUp(ctx context.Context, req domain.SignupRequest) (domain.SignupResponse, error) {

	existingUser, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return domain.SignupResponse{}, fmt.Errorf("error checking existing user: %w", err)
	}

	if existingUser != nil {
		return domain.SignupResponse{}, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.SignupResponse{}, fmt.Errorf("error hashing password: %w", err)
	}

	user := &domain.User{
		ID:        primitive.NewObjectID(),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Username:  req.Username,
		Password:  string(hashedPassword),
		Email:     req.Email,
		IsAdmin:   true,
	}

	if err := u.userRepo.CreateUser(ctx, user); err != nil {
		return domain.SignupResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	accessToken, err := u.authService.GenerateAccessToken(ctx, *user)
	if err != nil {
		return domain.SignupResponse{}, fmt.Errorf("error generating access token: %w", err)
	}

	_, err = u.authService.GenerateAndStoreRefreshToken(ctx, *user)
	if err != nil {
		return domain.SignupResponse{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	return domain.SignupResponse{
		AccessToken: accessToken,
	}, nil
}
