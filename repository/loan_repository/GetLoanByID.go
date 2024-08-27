package loan_repository

import (
	"context"
	"loan_tracker_api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *LoanRepository) GetLoanByID(ctx context.Context, id primitive.ObjectID) (*domain.Loan, error) {
	var loan domain.Loan
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&loan)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}
