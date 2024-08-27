package loan_usecase

import (
	"context"
	"loan_tracker_api/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *LoanUsecase) ViewAllLoans(ctx context.Context, filter domain.LoanStatus) ([]domain.LoanResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	query := bson.M{}
	if filter != "" && filter != "all" {
		query["status"] = filter
	}

	loans, err := u.loanRepo.GetLoans(ctx, query)
	if err != nil {
		return nil, err
	}

	var responses []domain.LoanResponse
	for _, loan := range loans {
		responses = append(responses, domain.LoanResponse{
			ID:           loan.ID,
			Amount:       loan.Amount,
			Term:         loan.Term,
			InterestRate: loan.InterestRate,
			Status:       loan.Status,
			CreatedAt:    loan.CreatedAt,
			UpdatedAt:    loan.UpdatedAt,
		})
	}

	return responses, nil
}
