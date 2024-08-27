package reset_token_repository

import (
	"loan_tracker_api/domain"
	"loan_tracker_api/mongo"
)

type resetTokenRepository struct {
	collection mongo.Collection
}

func NewResetTokenRepository(collection mongo.Collection) domain.ResetTokenRepository {
	return &resetTokenRepository{
		collection: collection,
	}
}
