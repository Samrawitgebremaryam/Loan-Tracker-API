package refresh_token_repository

import "loan_tracker_api/mongo"

type refreshTokenRepository struct {
	collection mongo.Collection
}

func NewRefreshTokenRepository(collection mongo.Collection) *refreshTokenRepository {
	return &refreshTokenRepository{
		collection: collection,
	}
}
