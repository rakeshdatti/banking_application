package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
	"github.com/rakesh/banking/app/mocks/service"
)



var mockAccountService *service.MockAccountService
var ah AccountHandler

func SetUpAccount(t *testing.T) func(){
	ctrl := gomock.NewController(t)
	mockAccountService = service.NewMockAccountService(ctrl)
	router = mux.NewRouter()
	return func(){
		defer ctrl.Finish()
	}
}

func Test_should_return_status_created_code_201(t *testing.T) {
	tearDown := SetUpAccount(t)
	defer tearDown()

	ah := AccountHandler{mockAccountService}
	request := dto.NewAccountRequest{
			CustomerId:  "1",
			AccountType: "saving",
			Amount: 5000.12,
	}
	reponse := &dto.NewAccountResponse{
		AccountID: "1234",
	}
	mockAccountService.EXPECT().NewAccount(request).Return(reponse,nil)

	account := ah.newAccount
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account",account)

	reqbodybytes ,_ := json.Marshal(request)
	req,_ :=http.NewRequest(http.MethodPost,"/customers/1/account",bytes.NewReader(reqbodybytes))

	//act
	recorder :=httptest.NewRecorder()
	router.ServeHTTP(recorder,req)

	//assert
	if recorder.Code!=http.StatusCreated{
		t.Error("Failed while testing the status code")
	}
	
}


func Test_should_return_error_when_creating_account_failed(t *testing.T) {
	tearDown := SetUpAccount(t)
	defer tearDown()

	ah := AccountHandler{mockAccountService}
	request := dto.NewAccountRequest{
			CustomerId:  "1",
			AccountType: "saving",
			Amount: 5000.12,
	}
	mockAccountService.EXPECT().NewAccount(request).Return(nil,errs.NewInternalServerError("some database error"))

	account := ah.newAccount
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account",account)

	reqbodybytes ,_ := json.Marshal(request)
	req,_ :=http.NewRequest(http.MethodPost,"/customers/1/account",bytes.NewReader(reqbodybytes))

	//act
	recorder :=httptest.NewRecorder()
	router.ServeHTTP(recorder,req)

	//assert
	if recorder.Code!=http.StatusInternalServerError{
		t.Error("Failed while testing the status code")
	}
	
}



func Test_should_return_status_code_200_sucessful_transaction(t *testing.T) {
	tearDown := SetUpAccount(t)
	defer tearDown()

	ah := AccountHandler{mockAccountService}
	requestbody := dto.NewTransactionRequest{
			Amount: 1000.12,
			TransactionType: "deposit",
	}
	reqestbytes,_ :=json.Marshal(requestbody)

	expectedRequest :=dto.NewTransactionRequest{
		CustomerId: "1",
		AccountId: "101",
		Amount: requestbody.Amount,
		TransactionType: requestbody.TransactionType,
	}


	reponse := &dto.NewTransactionResponse{
		TransactionId: "2342",
		AccountId: "101",
		Amount: 1000.12,
		TransactionType: "deposit",

	}
	mockAccountService.EXPECT().MakeTransaction(expectedRequest).Return(reponse,nil)

	
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}",ah.makeTransaction)
	req,_ :=http.NewRequest(http.MethodPost,"/customers/1/account/101",bytes.NewReader(reqestbytes))

	//act
	recorder :=httptest.NewRecorder()
	router.ServeHTTP(recorder,req)

	//assert
	if recorder.Code!=http.StatusOK{
		t.Error("Failed while testing the status code")
	}
	var actualResponse dto.NewTransactionResponse
	_ = json.NewDecoder(recorder.Body).Decode(&actualResponse)

	if actualResponse.TransactionId != "2342"{
		t.Errorf("Expected transaction ID txn123, but got %s", actualResponse.TransactionId)
	}
	
}


func Test_should_return_status_code_500_on_service_failure(t *testing.T) {
	tearDown := SetUpAccount(t)
	defer tearDown()

	ah := AccountHandler{mockAccountService}
	requestbody := dto.NewTransactionRequest{
			Amount: 1000.12,
			TransactionType: "deposit",
	}
	reqestbytes,_ :=json.Marshal(requestbody)

	expectedRequest :=dto.NewTransactionRequest{
		CustomerId: "1",
		AccountId: "101",
		Amount: requestbody.Amount,
		TransactionType: requestbody.TransactionType,
	}

	mockAccountService.EXPECT().MakeTransaction(expectedRequest).Return(nil,errs.NewInternalServerError("some database error"))

	
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}",ah.makeTransaction)
	req,_ :=http.NewRequest(http.MethodPost,"/customers/1/account/101",bytes.NewReader(reqestbytes))

	//act
	recorder :=httptest.NewRecorder()
	router.ServeHTTP(recorder,req)

	//assert
	if recorder.Code!=http.StatusInternalServerError{
		t.Error("Excepted code is 500 but got ",recorder.Code)
	}
	
}
