package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanStatus string

const (
	LoanStatusPending  LoanStatus = "pending"
	LoanStatusApproved LoanStatus = "approved"
	LoanStatusRejected LoanStatus = "rejected"
)

type Loan struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount       float64            `json:"amount" bson:"amount"`
	Term         int                `json:"term" bson:"term"` // Term in months or years, depending on your application
	InterestRate float64            `json:"interest_rate" bson:"interest_rate"`
	Status       LoanStatus         `json:"status" bson:"status"` // Pending, Approved, Rejected
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type LoanApplication struct {
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount       float64            `json:"amount" bson:"amount" binding:"required"`
	Term         int                `json:"term" bson:"term" binding:"required"`
	InterestRate float64            `json:"interest_rate" bson:"interest_rate" binding:"required"`
}

type LoanResponse struct {
	ID           primitive.ObjectID `json:"id"`
	Amount       float64            `json:"amount"`
	Term         int                `json:"term"`
	InterestRate float64            `json:"interest_rate"`
	Status       LoanStatus         `json:"status"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

type UpdateLoanStatusRequest struct {
	Status LoanStatus `json:"status" binding:"required"`
}

type LoanRepository interface {
	ApplyForLoan(ctx context.Context, loan *Loan) error
	GetLoanByID(ctx context.Context, id primitive.ObjectID) (*Loan, error)
	GetLoans(ctx context.Context, filter bson.M) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, id primitive.ObjectID, status LoanStatus) error
	DeleteLoan(ctx context.Context, id primitive.ObjectID) error
}

type LoanUsecase interface {
	ApplyForLoan(ctx context.Context, userID primitive.ObjectID, req LoanApplication) (LoanResponse, error)
	ViewLoanStatus(ctx context.Context, userID, loanID primitive.ObjectID) (LoanResponse, error)
	ViewAllLoans(ctx context.Context, filter LoanStatus) ([]LoanResponse, error)
	UpdateLoanStatus(ctx context.Context, loanID primitive.ObjectID, status LoanStatus) error
	DeleteLoan(ctx context.Context, loanID primitive.ObjectID) error
}
