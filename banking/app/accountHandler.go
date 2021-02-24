package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moisotico/banking/dto"
	"github.com/moisotico/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

// receive the request from a client, once received decode
func (h *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	//retrieve Id
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		// pass the request to the new account
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			// error
			writeResponse(w, appError.Code, appError.Message)
		} else {
			// status code is 201
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

// For transactions
// Format: /customers/2000/accounts/90720
func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get account_id & customer_id from URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	//decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		// build requested obj
		request.AccountId = accountId
		request.CustomerId = customerId

		// make the transaction
		account, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}

	}

}
