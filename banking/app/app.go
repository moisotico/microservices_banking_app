package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/moisotico/banking/domain"
	"github.com/moisotico/banking/service"
)

func sanityCheck() {
	// Variables to be checked on the sanityCheck
	keys := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASSWD", "DB_ADDR", "DB_PORT", "DB_NAME"}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Fatalf("Environment variable %s is not defined", k)
		}
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// defining routes using mux, for individual values, takes only numerical values
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// example on new customer

	//starting server: hardcoded direction
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
