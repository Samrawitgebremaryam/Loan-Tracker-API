package loan_usecase

import (
	"context"
	"loan_tracker_api/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecase) UpdateLoanStatus(ctx context.Context, loanID primitive.ObjectID, status domain.LoanStatus) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.loanRepo.UpdateLoanStatus(ctx, loanID, status)
}
