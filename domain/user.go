package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	IsAdmin   bool               `json:"isAdmin" bson:"isAdmin"`
}

type UserUsecase interface {
	SignUp(ctx context.Context, req SignupRequest) (SignupResponse, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	RequestPasswordReset(ctx context.Context, email, frontendBaseURL string) error
	ResetPassword(ctx context.Context, req ResetPasswordRequest) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByUsernameOrEmail(ctx context.Context, identifier string) (*User, error)
	UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}
