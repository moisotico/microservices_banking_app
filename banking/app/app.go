package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/moisotico/banking/domain"
	"github.com/moisotico/banking/service"
)

func sanityCheck() {
	// Variables to be checked on the sanityCheck
	keys := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range keys {
		if os.Getenv(k) == "" {
			log.Fatalf("Environment variable %s is not defined, Terminating application...", k)
		}
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	//wiring
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// defining routes using mux, for individual values, takes only numerical values
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// example on new customer

	//starting server: hardcoded direction
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {

	// use the env variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// get variables from environment variable
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)

	// defining the client hardcoded
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
