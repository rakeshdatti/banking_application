package domain

//usign sqlx lib 
import (
	"database/sql"
	// "log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rakesh/banking/app/errs"
	"github.com/rakesh/banking/app/logger"
)
type CustomerRepositorydb struct{
	client *sql.DB
}

func (d CustomerRepositorydb) FindAll(status string) ([]Customer,*errs.AppError){
	
	var rows *sql.Rows
	var err error
	if status =="" {
		findAllsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		rows,err = d.client.Query(findAllsql)
	}else{
		findAllsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where status=?"
		rows,err = d.client.Query(findAllsql,status)
	}
	
	if err!=nil{
		// log.Println("Error while query customer table "+ err.Error())
		logger.Error("Error while query customer table "+ err.Error())
		return nil,errs.NewInternalServerError("Unexpected database error")
	}
	defer rows.Close()
	 customers := make([]Customer,0)
	for rows.Next(){
		var c Customer
		err:= rows.Scan(&c.Id,&c.Name,&c.City,&c.ZipCode,&c.DateOfBirth,&c.Status)
		if err!=nil{
				// log.Println("Error while scanning customer"+err.Error())
				logger.Error("Error while scanning customer"+err.Error())
				return nil,errs.NewInternalServerError("Unexpected database error")
			// log.Println(" "+ err.Error())
		}
		customers =append(customers, c)
	}
	return customers,nil
}

func (d CustomerRepositorydb) FindbyId(id string) (*Customer,*errs.AppError){
	findbyidsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?" 
	row := d.client.QueryRow(findbyidsql,id)
	var c Customer
	err := row.Scan(&c.Id,&c.Name,&c.City,&c.ZipCode,&c.DateOfBirth,&c.Status)
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

func NewCustomerRepositorydb() CustomerRepositorydb{
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)   //maximum time a connection can be reused
	client.SetMaxOpenConns(10)  //maximum limit open connections
	client.SetMaxIdleConns(10) // idle connections always be avaiable in connection pool

	return CustomerRepositorydb{client}
}
