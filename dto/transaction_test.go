package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	//Arrange
	request := NewTransactionRequest{
		TransactionType: "invalid transaction type",
	}
	//Act
	appError := request.Validate()
	//Assert
	if appError.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Invalid message while testing the transaction type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invaid code while testin the trasaction type")
	}
}



func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	//arramge
	request := NewTransactionRequest{TransactionType: DEPOSIT,Amount: -100}
	//act
	appError := request.Validate()
	//Assert
	if appError.Message!="Amount connot less than zero"{
		t.Error("Invalid message while validating the amount")
	}

	if appError.Code!=http.StatusUnprocessableEntity{
		t.Error("Invalid code while validating the amount")
	}

}
