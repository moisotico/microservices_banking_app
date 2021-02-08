package domain

import "github.com/moisotico/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Interface for dbs and mock implementations
type CustomerRepository interface {
	// All customers
	FindAll() ([]Customer, *errs.AppError)
	// pointer in case we need nil
	ById(string) (*Customer, *errs.AppError)
}
