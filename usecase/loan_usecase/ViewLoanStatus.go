package loan_usecase

import (
	"context"
	"errors"
	"loan_tracker_api/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecase) ViewLoanStatus(ctx context.Context, userID, loanID primitive.ObjectID) (domain.LoanResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	loan, err := u.loanRepo.GetLoanByID(ctx, loanID)
	if err != nil {
		return domain.LoanResponse{}, err
	}

	if loan.UserID != userID {
		return domain.LoanResponse{}, errors.New("unauthorized access to loan")
	}

	return domain.LoanResponse{
		ID:           loan.ID,
		Amount:       loan.Amount,
		Term:         loan.Term,
		InterestRate: loan.InterestRate,
		Status:       loan.Status,
		CreatedAt:    loan.CreatedAt,
		UpdatedAt:    loan.UpdatedAt,
	}, nil
}
