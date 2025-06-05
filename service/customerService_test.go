package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	realdomain "github.com/rakesh/banking/app/domain"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/mocks/domain"
)

var mockrepoCustomer *domain.MockCustomerRepository

func SetupCustomer(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockrepoCustomer = domain.NewMockCustomerRepository(ctrl)
	return func() {
		ctrl.Finish()
	}

}
func Test_should_return_a_customer_with_status_code_200(t *testing.T) {
	tearDown := SetupCustomer(t)
	defer tearDown()

	service := NewCustomerService(mockrepoCustomer)
	
	_ = []realdomain.Customer{
		{Id: "1",Name: "rakesh",City: "rajam",DateOfBirth: "05/07/2002",ZipCode:"532122,",Status: "1"},
		{Id: "2",Name: "rocky",City: "rajam",DateOfBirth: "01/01/2002",ZipCode:"532122",Status: "0"},
	} 
	expectedReponse := []dto.CustomerResponse{
		{Id: "1",Name: "rakesh",City: "rajam",DateOfBirth: "05/07/2002",ZipCode:"532122,",Status: "1"},
	}
	var domainResponse []realdomain.Customer
	for _,dtocust := range expectedReponse{
		domainResponse =append(domainResponse,realdomain.Customer{
			Id: dtocust.Id,
			Name: dtocust.Name,
			City: dtocust.City,
			ZipCode: dtocust.City,
			DateOfBirth: dtocust.DateOfBirth,
		})
	}
	mockrepoCustomer.EXPECT().FindAll("1").Return(domainResponse,nil)
	//act 
	response,appError := service.GetAllCustomers("active")
	
	//assert
	if appError!=nil{
		t.Error("Error while retrive the customers from server")
	}
	if response[0].Name!="rakesh" && response[0].Status!="1"{
		t.Errorf("Expected first customer name to be 'rakesh', but got %s", response[0].Name)
	}

}
