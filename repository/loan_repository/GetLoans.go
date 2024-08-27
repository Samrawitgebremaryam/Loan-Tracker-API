package loan_repository

import (
	"context"
	"loan_tracker_api/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *LoanRepository) GetLoans(ctx context.Context, filter bson.M) ([]domain.Loan, error) {
	var loans []domain.Loan
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loans, nil
}
