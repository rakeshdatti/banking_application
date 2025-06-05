package domain

import (
	"github.com/rakesh/banking/app/dto"
	"github.com/rakesh/banking/app/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}


func (c *Customer) statusAsText() string{
	statusAsText := "active"
	if c.Status== "0"{
		statusAsText="inactive"
	}
	return statusAsText
}
func (c *Customer) ToDto() dto.CustomerResponse{
	
	return dto.CustomerResponse{
		Id: c.Id,
		Name:c.Name,
		City:c.City,
		ZipCode: c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status: c.statusAsText(),
	}
}

//go:generate mockgen -destination=../mocks/domain/mockCustomerRepository.go -package=domain github.com/rakesh/banking/app/domain CustomerRepository

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindbyId(string) (*Customer, *errs.AppError) //why pointer means,it customer is not there it will return nil,only pointers we can do.
}
