package loan_repository

import (
	"context"
	"loan_tracker_api/domain"
)

func (r *LoanRepository) ApplyForLoan(ctx context.Context, loan *domain.Loan) error {
	_, err := r.collection.InsertOne(ctx, loan)
	return err
}
