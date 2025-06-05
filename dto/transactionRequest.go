package dto

import "github.com/rakesh/banking/app/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type NewTransactionRequest struct {
	CustomerId      string  `json: "-"`
	Amount          float64 `json: "amount"`
	AccountId       string  `json: "account_id"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL 
}
func (t NewTransactionRequest) IsTransactionTypeDeposit() bool {
	return t.TransactionType == DEPOSIT 
}

func (t NewTransactionRequest) Validate() *errs.AppError {
	if !t.IsTransactionTypeWithdrawal() && !t.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if t.Amount < 0 {
		return errs.NewValidationError("Amount connot less than zero")
	}
	return nil
}

