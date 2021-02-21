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
