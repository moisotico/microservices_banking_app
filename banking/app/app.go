package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// using gorilla/mux
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// defining routes using mux
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// example on new customer

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
