package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	realdomain "github.com/rakesh/banking/app/domain"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
	"github.com/rakesh/banking/app/mocks/domain"
)

var mockrepo *domain.MockAccountRepository

func SetupAccount(t *testing.T) func(){
	ctrl :=gomock.NewController(t)
	mockrepo =domain.NewMockAccountRepository(ctrl)
	return func(){
		ctrl.Finish()
	}
	
}
func Test_should_return_a_validation_error_response_when_request_is_not_validated(t *testing.T) {
	//arrange
	request := dto.NewAccountRequest{
		CustomerId: "1",
		AccountType: "saving",
		Amount: 0,
	}
	service := NewAccountService(nil)

	//act 
	_, appError := service.NewAccount(request)
	if appError==nil{
		t.Error("error while testing the   new account validation")
	}

}


func Test_should_return_a_error_from_server_side_if_the_new_account_cannot_be_Created(t *testing.T){		
	tearDown := SetupAccount(t)
	defer tearDown()
	service := NewAccountService(mockrepo)

	//act 
	req := dto.NewAccountRequest{
		CustomerId: "1",
		AccountType: "saving",
		Amount: 6000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId, 
		AccountType: req.AccountType,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount:      req.Amount,
		Status:      "1",
	}
	mockrepo.EXPECT().Save(account).Return(nil,errs.NewInternalServerError("Unexptected database error"))

	//assert
	_,appErr :=service.NewAccount(req)
	if appErr==nil{
		t.Error("test failed while validating error for creating  a new account")
	}
}


func Test_should_return_a_account_from_server_side_if_the_new_account_created_sucesfully(t *testing.T){		
	tearDown := SetupAccount(t)
	defer tearDown()
	service := NewAccountService(mockrepo)
	//act 
	req := dto.NewAccountRequest{
		CustomerId: "1",
		AccountType: "saving",
		Amount: 6000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId, 
		AccountType: req.AccountType,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount:      req.Amount,
		Status:      "1",
	}
	accountwithid := account
	accountwithid.AccountID="201"
	mockrepo.EXPECT().Save(account).Return(&accountwithid,nil)
	//act
	newAccount,appErr :=service.NewAccount(req)
	if appErr!=nil{
		t.Error("test failed while validating error for creating  a new account")
	}
	if newAccount.AccountID !=accountwithid.AccountID{
		t.Error("Failed while matching accoundid")
	}
}


// func Test_should_return_a_account_from_server_side_if_the_new_account_created_sucesfully(t *testing.T){		
// 	ctrl :=gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockrepo :=domain.NewMockAccountRepository(ctrl)
// 	service := NewAccountService(mockrepo)

// 	//act 
// 	req := dto.NewAccountRequest{
// 		CustomerId: "1",
// 		AccountType: "saving",
// 		Amount: 6000,
// 	}
// 	account := realdomain.Account{
// 		CustomerId:  req.CustomerId, 
// 		AccountType: req.AccountType,
// 		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
// 		Amount:      req.Amount,
// 		Status:      "1",
// 	}
// 	mockrepo.EXPECT().Save(account).Return(nil,errs.NewInternalServerError("Unexptected database error"))

// 	//act
// 	_,appErr :=service.NewAccount(req)
// 	if appErr==nil{
// 		t.Error("test failed while validating error for creating  a new account")
// 	}
// }


//For Transaction testing 

func Test_should_return_a_validation_error_response_when_request_is_not_validated_for_transaction(t *testing.T) {
	//arrange
	request := dto.NewTransactionRequest{
		CustomerId: "1",
		TransactionType: "invalidtype",
		Amount: -1,
		AccountId: "101",
	}
	service := NewAccountService(nil)

	//act 
	_, appError := service.MakeTransaction(request)
	if appError==nil{
		t.Error("error while testing the   new account validation")
	}

}


func  TestMakeTransaction_SuccessfulDeposit(t *testing.T){
		tearDown := SetupAccount(t)
		defer tearDown()
		service := NewAccountService(mockrepo)

		//act
		req := dto.NewTransactionRequest{
			TransactionType: "deposit",
			Amount: 1000.0,
			AccountId: "101",
		}
		domainRequest := realdomain.Transaction{
			Amount:      req.Amount,
			AccountId: req.AccountId,
			TransactionType: req.TransactionType,
			TransactionDate:  time.Now().Format("2006-01-02 15:04:05"),
		}
		expectedTransaction := &realdomain.Transaction{
			TransactionId: "101",
			AccountId:       "123",
			Amount:          1000,
			TransactionType: "deposit",
			TransactionDate:  time.Now().Format("2006-01-02 15:04:05"),
		}
		mockrepo.EXPECT().SaveTransaction(domainRequest).Return(expectedTransaction, nil)

		//assert
		transactionreponse ,appErr := service.MakeTransaction(req)

		if appErr!=nil{
			t.Error("test failed while deposit the transaction")
		}
		if transactionreponse.TransactionId!=expectedTransaction.TransactionId{
			t.Error("Failed while matching transactionID")
		}

}


func  Test_shuould_return_errro_from_server_failed_transaction_Deposit(t *testing.T){
	tearDown := SetupAccount(t)
	defer tearDown()
	service := NewAccountService(mockrepo)

	//act
	req := dto.NewTransactionRequest{
		TransactionType: "deposit",
		Amount: 1000.0,
		AccountId: "101",
	}
	domainRequest := realdomain.Transaction{
		Amount:      req.Amount,
		AccountId: req.AccountId,
		TransactionType: req.TransactionType,
		TransactionDate:  time.Now().Format("2006-01-02 15:04:05"),
	}
	mockrepo.EXPECT().SaveTransaction(domainRequest).Return(nil, errs.NewInternalServerError("some database error"))

	//assert
	_ ,appErr := service.MakeTransaction(req)

	if appErr==nil{
		t.Error("test failed while deposit the transaction")
	}

}



func  Test_MakeTransaction_Successful_withdrawal(t *testing.T){
	tearDown := SetupAccount(t)
	defer tearDown()

	service := NewAccountService(mockrepo)
	req := dto.NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount: 1000.0,
		AccountId: "101",
	}
	domainRequest := realdomain.Transaction{
		Amount:      req.Amount,
		AccountId: req.AccountId,
		TransactionType: req.TransactionType,
		TransactionDate:  time.Now().Format("2006-01-02 15:04:05"),
	}

	mockAccount := &realdomain.Account{
		Amount: 10000.0,
		AccountID: "123",
	}
	expectedTransaction := &realdomain.Transaction{
		TransactionId: "101",
		AccountId:       req.AccountId,
		Amount:          9000.0,
		TransactionType: req.TransactionType,
		TransactionDate:  req.TransactionDate,
	}
	mockrepo.EXPECT().FindById(req.AccountId).Return(mockAccount,nil)
	mockrepo.EXPECT().SaveTransaction(domainRequest).Return(expectedTransaction,nil)

	//act 
	transactionResponse ,appError:=service.MakeTransaction(req)
	if appError!=nil{
		t.Errorf("Expected no error, but got %v", appError)
	}
	if transactionResponse ==nil{
		t.Error("Expected Response ,but got nil")
	}
	if transactionResponse.TransactionId != expectedTransaction.TransactionId{
		t.Errorf("Expected Transaction %v,but got %v",expectedTransaction.TransactionId,transactionResponse.TransactionId)
	}
}

func  Test_should_return_error_failure_of_withdrwal_transaction(t *testing.T){
	tearDown := SetupAccount(t)
	defer tearDown()

	service := NewAccountService(mockrepo)
	req := dto.NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount: 1000.0,
		AccountId: "101",
	}
	domainRequest := realdomain.Transaction{
		Amount:      req.Amount,
		AccountId: req.AccountId,
		TransactionType: req.TransactionType,
		TransactionDate:  time.Now().Format("2006-01-02 15:04:05"),
	}

	mockAccount := &realdomain.Account{
		Amount: 2000.0,
		AccountID: "123",
	}
	
	mockrepo.EXPECT().FindById(req.AccountId).Return(mockAccount,nil)
	mockrepo.EXPECT().SaveTransaction(domainRequest).Return(nil,errs.NewInternalServerError("some database error"))

	//act 
	_ ,appError:=service.MakeTransaction(req)
	if appError==nil{
		t.Errorf("Expected  error, but got %v", appError)
	}

}




func  Test_should_return_error_failure_of_withdrawal_amount_greatere_than_balance_amount(t *testing.T){
	tearDown := SetupAccount(t)
	defer tearDown()

	service := NewAccountService(mockrepo)
	req := dto.NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount: 1000.0,
		AccountId: "101",
	}

	mockAccount := &realdomain.Account{
		Amount: 200.0,
		AccountID: "123",
	}
	
	mockrepo.EXPECT().FindById(req.AccountId).Return(mockAccount,nil)
	//act 
	_ ,appError:=service.MakeTransaction(req)

	//assert
	if appError==nil{
		t.Error("Expected validation error for insufficient balance, but got nil")
	}
	if appError==nil && appError.Message != "Insufficient balance in the account"{
		t.Errorf("Expected Insufficient balance in the account,but got %v",appError.Message)
	}

}
