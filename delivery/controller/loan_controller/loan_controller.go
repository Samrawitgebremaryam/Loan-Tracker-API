package loan_controller

import (
	"loan_tracker_api/domain"
)

type LoanController struct {
	loanUsecase domain.LoanUsecase
}

func NewLoanController(loanUsecase domain.LoanUsecase) *LoanController {
	return &LoanController{
		loanUsecase: loanUsecase,
	}
}
