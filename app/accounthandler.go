package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/service"
)

type AccountHandler struct {
	service service.AccountService
}


func (ah *AccountHandler) newAccount(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err!=nil{
		writeResponse(w,http.StatusBadGateway,err.Error())
	}else{
		request.CustomerId=customerId
		account, appError :=ah.service.NewAccount(request)
		if appError!=nil{
			writeResponse(w,appError.Code,appError.AsMessage())
		}else{
			writeResponse(w,http.StatusCreated,account)
		}
	}
}





func (ah *AccountHandler) makeTransaction(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerID :=vars["customer_id"]
	accountID :=vars["account_id"]
	var request dto.NewTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err !=nil{
		writeResponse(w,http.StatusBadRequest,err.Error())
	}else{
		request.CustomerId=customerID
		request.AccountId=accountID
		account ,appError := ah.service.MakeTransaction(request)
		if appError!=nil{
			writeResponse(w,appError.Code,appError.AsMessage())
		}else{
			writeResponse(w,http.StatusOK,account)
		}
	}

}
