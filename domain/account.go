package domain

import (
	"time"

	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
)

type Account struct {
	AccountID   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string `db:"status"`
}

func (a *Account) ToNewAccountReponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountID: a.AccountID,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/rakesh/banking/app/domain AccountRepository

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(string) (*Account,*errs.AppError)
	SaveTransaction(Transaction)(*Transaction,*errs.AppError)
}


func (a *Account) CanWithdraw(amount float64)bool{
	if a.Amount>amount{
		return true
	}
	return false
}


func NewAccount(customerId,accountType string,amount float64) Account{
	return Account{
		CustomerId:  customerId,
		AccountType: accountType,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount:      amount,
		Status:      "1",
	}
}
