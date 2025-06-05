package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rakesh/banking/app/errs"
	"github.com/rakesh/banking/app/logger"
)
type CustomerRepositorydbsqlx struct{
	client *sqlx.DB
}

func (d CustomerRepositorydbsqlx) FindAll(status string) ([]Customer,*errs.AppError){
	
	// var rows *sql.Rows
	var err error
	customers := make([]Customer,0)
	if status =="" {
		findAllsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err = d.client.Select(&customers,findAllsql)	
	}else{
		findAllsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where status=?"
		err =d.client.Select(&customers,findAllsql,status)
	}
	
	if err!=nil{
		// log.Println("Error while query customer table "+ err.Error())
		logger.Error("Error while query customer table "+ err.Error())
		return nil,errs.NewInternalServerError("Unexpected database error")
	}
	// defer rows.Close()

	return customers,nil
}

func (d CustomerRepositorydbsqlx) FindbyId(id string) (*Customer,*errs.AppError){
	findbyidsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?" 
	var c Customer
	err := d.client.Get(&c,findbyidsql,id)
	if err!=nil{
		if err==sql.ErrNoRows{
			return nil,errs.NewNotFoundError("Customer not found")
		}else{
			// log.Println("Error while scanning customer"+err.Error())
			logger.Error("Error while scanning customer"+err.Error())
			return nil,errs.NewInternalServerError("Unexpected database error")
		}
	}
	return &c,nil
}

func NewCustomerRepositorydbsqlx(dbClient *sqlx.DB) CustomerRepositorydbsqlx{

	return CustomerRepositorydbsqlx{dbClient}
}
