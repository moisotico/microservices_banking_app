package domain

import "github.com/moisotico/banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// Interface for dbs and mock implementations
type CustomerRepository interface {
	// status can be active (1) or inactive (0)

	// All customers
	FindAll(status string) ([]Customer, *errs.AppError)
	// pointer in case we need nil
	ById(string) (*Customer, *errs.AppError)
}
