package loan_repository

import (
	"loan_tracker_api/domain"

	"loan_tracker_api/mongo"
)

type LoanRepository struct {
	collection mongo.Collection
}

func NewLoanRepository(collection mongo.Collection) domain.LoanRepository {
	return &LoanRepository{
		collection: collection,
	}
}
