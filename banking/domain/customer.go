package domain

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
	FindAll() ([]Customer, error)
}
