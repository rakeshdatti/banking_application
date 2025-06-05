package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
	"github.com/rakesh/banking/app/mocks/service"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService


func SetUpCustomer(t *testing.T) func(){
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch =CustomerHandlers{mockService}
	router = mux.NewRouter()
	return func(){
		defer ctrl.Finish()
	}
}

func Test_should_Return_customers_with_status_code_200(t *testing.T) {
	//arrange 
	tearDown := SetUpCustomer(t)
	defer tearDown()
	// ctrl := gomock.NewController(t)

	//it is importatn if you don;t call finish it is not going to test your interactions
	// defer ctrl.Finish()
	
	//service.NewMockCustomerService(t) takes controller and controller manages the state for this mock
	// mockService := service.NewMockCustomerService(ctrl)
	dummyCustomers := []dto.CustomerResponse{
		{Id: "1",Name: "rakesh",City: "rajam",ZipCode: "532122",DateOfBirth: "05/07/2002",Status: "1"},
		{Id: "2",Name: "rocky",City: "rajam",ZipCode: "532122",DateOfBirth: "01/01/2001",Status: "0"},
	}
	mockService.EXPECT().GetAllCustomers("").Return(dummyCustomers,nil)
	// ch =CustomerHandlers{mockService}

	// router := mux.NewRouter()
	 router.HandleFunc("/customers",ch.getAllCustomers)
	//when call is made to the customer enpoint getAllCusotomers ,insdie this instead of calling default custoemr service 
	//it will call to mock customer service  GetallCustomerMEthod

	//6.http request 
	request ,_ :=http.NewRequest(http.MethodGet,"/customers",nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder,request)

	//Assert
	if recorder.Code!=http.StatusOK{
		t.Error("Failed while testing the status code")
	}
}


func Test_should_Return_status_code_500_with_error_message(t *testing.T) {
	//arrange 
	tearDown := SetUpCustomer(t)
	defer tearDown()
	mockService.EXPECT().GetAllCustomers("").Return(nil,errs.NewInternalServerError("some database error"))
	router.HandleFunc("/customers",ch.getAllCustomers)
	request ,_ :=http.NewRequest(http.MethodGet,"/customers",nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder,request)

	//Assert
	if recorder.Code!=http.StatusInternalServerError{
		t.Error("Failed while testing the status code")
	}
}

// func Test_should_Return_status_code_500_with_error_message(t *testing.T) {
// 	//arrange 
// 	ctrl := gomock.NewController(t)
// 	//it is importatn if you don;t call finish it is not going to test your interactions
// 	defer ctrl.Finish()
	
// 	//service.NewMockCustomerService(t) takes controller and controller manages the state for this mock
// 	mockService := service.NewMockCustomerService(ctrl)
// 	mockService.EXPECT().GetAllCustomers("").Return(nil,errs.NewInternalServerError("some database error"))
// 	ch :=CustomerHandlers{mockService}

// 	router := mux.NewRouter()
// 	router.HandleFunc("/customers",ch.getAllCustomers)
// 	//when call is made to the customer enpoint getAllCusotomers ,insdie this instead of calling default custoemr service 
// 	//it will call to mock customer service  GetallCustomerMEthod

// 	//6.http request 
// 	request ,_ :=http.NewRequest(http.MethodGet,"/customers",nil)

// 	//Act
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder,request)

// 	//Assert
// 	if recorder.Code!=http.StatusInternalServerError{
// 		t.Error("Failed while testing the status code")
// 	}
// }



func Test_should_Return_customer_by_id_with_status_code_200(t *testing.T) {
	//arrange 
	tearDown := SetUpCustomer(t)
	defer tearDown()
	dummycustomer := &dto.CustomerResponse{
		Id: "1",
		Name: "rakesh",
		City: "rajam",
		ZipCode: "532122",
		DateOfBirth: "05/07/2002",
		Status: "1",
	}

	mockService.EXPECT().GetCustomersbyId("1").Return(dummycustomer,nil)
	router.HandleFunc("/customers/{customer_id:[0-9]+}",ch.getCustomerbyId)

	request ,_ :=http.NewRequest(http.MethodGet,"/customers/1",nil)

	//act 
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder,request)

	//assert
	if recorder.Code !=http.StatusOK{
		t.Error("Failed while testing the status code")
	}

}


func Test_should_Return_status_500_with_error_customer_by_id(t *testing.T) {
	tearDown := SetUpCustomer(t)
	defer tearDown()

	mockService.EXPECT().GetCustomersbyId("1").Return(nil,errs.NewInternalServerError("some database error"))
	router.HandleFunc("/customers/{customer_id:[0-9]+}",ch.getCustomerbyId)

	request ,_ :=http.NewRequest(http.MethodGet,"/customers/1",nil)

	//act 
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder,request)

	//assert
	if recorder.Code !=http.StatusInternalServerError{
		t.Error("Failed while testing the status code")
	}

}
