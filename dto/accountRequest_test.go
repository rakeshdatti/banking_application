package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_createing_account_the_amount_less_than_basic_amount(t *testing.T) {
	//Arrange
	request := NewAccountRequest{
		Amount: 3000,
		AccountType: "saving",
	}
	//Act
	appError := request.Validate()
	//Assert
	if appError.Message != "Amount should at least 5000 to create a account" {
		t.Error("Invalid message while testing the creating account")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invaid code while testing the create account")
	}
}



func Test_should_return_error_when_creating_account_the_account_type_is_not_saving_or_checking(t *testing.T) {
	//Arrange
	request := NewAccountRequest{
		AccountType: "invalidtype",
		Amount: 5000,
	}
	//Act
	appError := request.Validate()
	//Assert
	if appError.Message != "Account should be type checking or saving"{
		t.Error("Invalid message while testing the creating account")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invaid code while testing the create account")
	}
}
