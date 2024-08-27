package user_repository

import (
	"loan_tracker_api/domain"
	"loan_tracker_api/mongo"
)

type userRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) domain.UserRepository {
	return &userRepository{
		collection: collection,
	}
}
