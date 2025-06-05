package app

import (
	"encoding/json"
	"log"

	// "encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakesh/banking/app/logger"
	"github.com/rakesh/banking/app/service"
)

// type Customer struct {
// 	Name    string `json: "name" xml" "name"`
// 	City    string `json: "city" xml: "city"`
// 	ZipCode string `json: "zipcode" xml:"zipcode"`
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello world")
// }

type CustomerHandlers struct {
	service service.CustomerService
}



func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"rakesh", "vizag", "532445"},
	// 	{"rocky", "vizag", "532445"},
	// }
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err!=nil{
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		logger.Info("retriving the customers data")
		writeResponse(w,http.StatusOK,customers)
	}



	//by default header will return json ,if you want xml we need say
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }











	//to make the type as json/application
	// w.Header().Add("Content-Type","application/json")
	//marshling the data to json representation
	// json.NewEncoder(w).Encode(customers)

	//marshling data strutures to XML representation
	// w.Header().Add("Content-Type","application/xml")
	// xml.NewEncoder(w).Encode(customers)

}

func (ch *CustomerHandlers) getCustomerbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomersbyId(id)
	if err != nil {
		writeResponse(w,err.Code,err.AsMessage())
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err.Code)
		// json.NewEncoder(w).Encode(err.AsMessage())
		// fmt.Fprintf(w,err.Message)
	} else {
		writeResponse(w,http.StatusOK,customer)
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(customer)
	}
}



func writeResponse(w http.ResponseWriter,code int,data interface{}){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data);err!=nil{
		// panic(err)
		log.Panicln("Error in Encoding response",err)
		http.Error(w,"Internal Server Error ",http.StatusInternalServerError)
	}
}




// func getCustomer(w http.ResponseWriter, r *http.Request) {

// 	//you need to create map of route variables
// 	vars := mux.Vars(r)

// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "post request received")
// }
