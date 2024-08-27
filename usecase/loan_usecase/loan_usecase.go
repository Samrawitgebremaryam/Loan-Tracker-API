package loan_usecase

import (
	"loan_tracker_api/domain"
	"time"
)

type LoanUsecase struct {
	loanRepo       domain.LoanRepository
	contextTimeout time.Duration
}

func NewLoanUsecase(loanRepo domain.LoanRepository, timeout time.Duration) domain.LoanUsecase {
	return &LoanUsecase{
		loanRepo:       loanRepo,
		contextTimeout: timeout,
	}
}
