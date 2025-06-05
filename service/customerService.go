package service

import (
	"github.com/rakesh/banking/app/domain"
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/rakesh/banking/app/service CustomerService
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse,*errs.AppError)
	GetCustomersbyId(string) (*dto.CustomerResponse,*errs.AppError)
}


type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse,*errs.AppError){
	if status== "active"{
		status=  "1"
	}else if status=="inactive"{
		status= "0"
	}else{
		status=""
	}
	c,err := s.repo.FindAll(status)
	if err!=nil{
		return nil,err 
	}

	var response []dto.CustomerResponse
	for _,c:= range c{
		resp := c.ToDto()
		response =append(response,resp)
	}
	
	return response,nil
}

func (s DefaultCustomerService) GetCustomersbyId(id string) (*dto.CustomerResponse,*errs.AppError){
	c,err := s.repo.FindbyId(id)
	if err!=nil{
		return nil,err
	}
	
	response :=c.ToDto()
	return &response,nil
}


func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repository}
}
