package service

import (
	"time"

	"github.com/rakesh/banking/app/domain"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
)


const dbTSLayout = "2006-01-02 15:04:05"


//go:generate mockgen -destination=../mocks/service/mockAccountService.go -package=service github.com/rakesh/banking/app/service AccountService
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.NewTransactionRequest)(*dto.NewTransactionResponse,*errs.AppError)
}



type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (a DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	//validate the request
	if err := req.Validate();err != nil {
		return nil, err
	}
	//we need convert the dto to domain object and pass to repo
	// da := domain.Account{
	// 	CustomerId:  req.CustomerId,
	// 	AccountType: req.AccountType,
	// 	OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
	// 	Amount:      req.Amount,
	// 	Status:      "1",
	// }
	da := domain.NewAccount(req.CustomerId,req.AccountType, req.Amount)
	if newAccount, err := a.repo.Save(da);err != nil {
		return nil, err
	}else{
		resp := newAccount.ToNewAccountReponseDto()
		return &resp, nil
	}
}


func (a DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest)(*dto.NewTransactionResponse,*errs.AppError){
	//incoming request validation
	if err := req.Validate(); err!=nil{
		return nil,err
	}
	//server side validation for checking the avaiable balance in the account
	if req.IsTransactionTypeWithdrawal(){
		account,err := a.repo.FindById(req.AccountId)
		if err!=nil{
			return nil,err 
		}
		if !account.CanWithdraw(req.Amount){
			return nil,errs.NewValidationError("Insufficient balance in the account")
		}
	}
	//if all is well ,build the domain object &save the transaction

	tobject := domain.Transaction{
		AccountId: req.AccountId,
		TransactionType: req.TransactionType,
		Amount: req.Amount,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction ,appError := a.repo.SaveTransaction(tobject)
	if appError!=nil{
		return nil,appError
	}
	response:=transaction.ToDto()
	return &response,nil

	
}



func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

