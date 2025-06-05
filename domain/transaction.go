package domain

import (
	"github.com/rakesh/banking/app/dto"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"stringaccount_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t *Transaction) ToDto() dto.NewTransactionResponse{
	return dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		AccountId: t.AccountId,
		Amount: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == "withdrawal" {
		return true
	}
	return false
}
