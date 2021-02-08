package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moisotico/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/* 	customers := []Customer{
		{Name: "Moises", City: "San Jose", Zipcode: "11505"},
		{Name: "Woises", City: "San Pedro", Zipcode: "11505"},
	} */

	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	// Do not change status code manually, might break
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

// Created function based on DRY principle
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
