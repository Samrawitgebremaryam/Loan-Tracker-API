package reset_token_repository

import (
	"context"
	"loan_tracker_api/domain"
)

func (r *resetTokenRepository) StoreResetToken(ctx context.Context, token domain.PasswordResetToken) error {
	_, err := r.collection.InsertOne(ctx, token)
	return err
}
