package loan_usecase

import (
	"context"
	"loan_tracker_api/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *LoanUsecase) ApplyForLoan(ctx context.Context, userID primitive.ObjectID, req domain.LoanApplication) (domain.LoanResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	loan := domain.Loan{
		ID:           primitive.NewObjectID(),
		UserID:       userID,
		Amount:       req.Amount,
		Term:         req.Term,
		InterestRate: req.InterestRate,
		Status:       domain.LoanStatusPending,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := u.loanRepo.ApplyForLoan(ctx, &loan)
	if err != nil {
		return domain.LoanResponse{}, err
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
