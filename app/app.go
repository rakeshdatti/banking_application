package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rakesh/banking/app/domain"
	"github.com/rakesh/banking/app/service"
)

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	dbClient := getdbClient()
	//wiring
	// customerhandlers := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// proper wiring
	crdbsqlx := domain.NewCustomerRepositorydbsqlx(dbClient)
	customerhandlers := CustomerHandlers{service: service.NewCustomerService(crdbsqlx)}

	//creating handler for account
	ardbsqlx := domain.NewAccountRepositorydbsqlx(dbClient)
	accounthandlers := AccountHandler{service: service.NewAccountService(ardbsqlx)}

	router.HandleFunc("/customers", customerhandlers.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerhandlers.getCustomerbyId)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accounthandlers.newAccount)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", accounthandlers.makeTransaction)

	// //routed
	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/getcustomers", getAllCustomers)
	// //starting server
	// fmt.Println("server starting at 8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// //creat own multiplexer
	//  mux := http.NewServeMux()
	// mux.HandleFunc("/greet", greet)
	// mux.HandleFunc("/getcustomers", getAllCustomers)
	// //starting server
	// fmt.Println("server starting at 8080")
	// log.Fatal(http.ListenAndServe(":8080", mux))

	// router := mux.NewRouter()

	// router.HandleFunc("/customers", getAllCustomers)

	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// router.HandleFunc("/customers", getAllCustomers)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	//starting server

	server_address := os.Getenv("SERVER_ADDRESS")
	server_port := os.Getenv("SERVER_PORT")
	fmt.Printf("server starting at %s:%s", server_address, server_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", server_address, server_port), router))

}

func getdbClient() *sqlx.DB {
	dbuser := os.Getenv("DB_USER")
	dbpasswd := os.Getenv("DB_PASSWD")
	dbaddress := os.Getenv("DB_ADDR")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	//the below code is creating a connection pool
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpasswd, dbaddress, dbport, dbname)

	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3) //maximum time a connection can be reused
	client.SetMaxOpenConns(10)                 //maximum limit open connections
	client.SetMaxIdleConns(10)                 // idle connections always be avaiable in connection pool
	return client
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" || os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWD") == "" || os.Getenv("DB_ADDR") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("environment variable is missing")
	}
}
