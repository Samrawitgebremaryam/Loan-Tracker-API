package loan_repository

import (
	"context"
	"loan_tracker_api/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LoanRepository) UpdateLoanStatus(ctx context.Context, id primitive.ObjectID, status domain.LoanStatus) error {
	update := bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
