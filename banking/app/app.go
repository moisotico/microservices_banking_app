package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moisotico/banking/domain"
	"github.com/moisotico/banking/service"
)

func Start() {
	// using gorilla/mux
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// defining routes using mux, for individual values, takes only numerical values
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// example on new customer

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
