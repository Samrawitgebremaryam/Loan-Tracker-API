package loan_usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecase) DeleteLoan(ctx context.Context, loanID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.loanRepo.DeleteLoan(ctx, loanID)
}
