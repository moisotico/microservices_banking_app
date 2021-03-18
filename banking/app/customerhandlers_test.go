package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/moisotico/banking/dto"
	"github.com/moisotico/banking/errs"
	"github.com/moisotico/banking/mocks/service"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()

	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	//Arange
	teardown := setup(t)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Moi", City: "San Pedro", Zipcode: "10105", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Woi", City: "San Pedro", Zipcode: "10105", DateOfBirth: "2000-01-01", Status: "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	// http request
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}

}

func Test_should_return_customers_with_status_code_500_with_error_message(t *testing.T) {

	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("a database error"))

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
