package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/moisotico/banking/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"

	// Quering the database
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Go trough all the rows & dereference via Scan
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customer table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")

		}
		customers = append(customers, c)
	}
	return customers, nil
}

// helper function to get customer by id
func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	// we get only one query
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// helper function to make only one connection to DB
func NewCustomerRepositoryDB() CustomerRepositoryDB {

	// defining the client
	client, err := sql.Open("mysql", "root:testpassword@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
