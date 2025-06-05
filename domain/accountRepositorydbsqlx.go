package domain

import (
	"strconv"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rakesh/banking/app/errs"
	"github.com/rakesh/banking/app/logger"
)



type AccountRepositorydbsqlx struct {
	client *sqlx.DB
}

func (d AccountRepositorydbsqlx) Save(a Account) (*Account, *errs.AppError){
	sqlxInsert := "INSERT INTO accounts(customer_id,opening_date,account_type,amount,status)values(?,?,?,?,?)"
	result,err :=d.client.Exec(sqlxInsert,a.CustomerId,a.OpeningDate,a.AccountType,a.Amount,a.Status)
	if err!=nil{
		logger.Error("Error while creating an account "+err.Error())
		return nil,errs.NewInternalServerError("Unexpected Error from Database")
	}
	id,err :=result.LastInsertId()
	if err!=nil{
		logger.Error("Error while getting id from  accounts "+err.Error())
		return nil,errs.NewInternalServerError("Unexpected Error from Database")
	}
	a.AccountID=strconv.FormatInt(id,10)
	return &a,nil 
} 



func (d AccountRepositorydbsqlx)FindById(accountId string)(*Account,*errs.AppError){
	findbyidsql := "select account_id,customer_id,opening_date,account_type,amount,status from accounts where account_id=?" 
	var a Account
	err := d.client.Get(&a,findbyidsql,accountId)
	if err!=nil{
		if err==sql.ErrNoRows{
			return nil,errs.NewNotFoundError("Customer not found")
		}else{
			// log.Println("Error while scanning customer"+err.Error())
			logger.Error("Error while scanning customer "+err.Error())
			return nil,errs.NewInternalServerError("Unexpected database error")
		}
	}
	return &a,nil
}


func(d AccountRepositorydbsqlx) SaveTransaction(t Transaction)(*Transaction,*errs.AppError){
	//starting with database transaction block 
	tx ,err :=d.client.Begin()
	if err!=nil{
		logger.Error("Errow while starting a new Transaction for bank account transaction"+err.Error())
		return nil,errs.NewInternalServerError("Unexpected database error")
	}
	result,err:=tx.Exec("INSERT INTO transactions (account_id,amount,transaction_type,transaction_date) values(?,?,?,?)",t.AccountId,t.Amount,t.TransactionType,t.TransactionDate)
		
	if err != nil {
		tx.Rollback()
		logger.Error("Error while inserting transaction: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected database error")
		}
		
	//updating account balance based on withdraw or deposit
	if t.IsWithdrawal(){
			_,err = tx.Exec("UPDATE accounts SET amount=amount-? where account_id=?",t.Amount,t.AccountId)
	}else{
		
		_,err = tx.Exec("UPDATE accounts SET amount=amount+? where account_id=?",t.Amount,t.AccountId)
	}
	
	//in case of error Roolback , and changed from the both tables will retrived
	if err!=nil{
		tx.Rollback()
		logger.Error("Error while saving the Transaction"+err.Error())
		return nil,errs.NewInternalServerError("Unexpted databse error")
	}
	//commit the transaction if all good
	err =tx.Commit()
	if err !=nil{
		logger.Error("Error while commiting  the Transaction for bank account"+err.Error())
		return nil,errs.NewInternalServerError("Unexpted databse error")
	}
	//getting the lasttransactionid from transaction table
	transactionId,err := result.LastInsertId()
	if err !=nil{
		logger.Error("Error while getting   the lasttransactionid"+err.Error())
		return nil,errs.NewInternalServerError("Unexpted databse error")
	}
	//getting latest account information from the accounts table
	account,appError := d.FindById(t.AccountId)
	if appError!=nil{
		return nil,appError
	}
	t.TransactionId=strconv.FormatInt(transactionId,10)

	//updating the transaction struct with latest balance
	t.Amount=account.Amount
	return &t,nil

}

func NewAccountRepositorydbsqlx(dbClient *sqlx.DB) AccountRepositorydbsqlx{
	return AccountRepositorydbsqlx{client: dbClient}
}
