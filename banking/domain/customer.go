package domain

import (
	"github.com/moisotico/banking/dto"
	"github.com/moisotico/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// Interface for dbs and mock implementations
type CustomerRepository interface {
	// status can be active (1) or inactive (0)

	// All customers
	FindAll(status string) ([]Customer, *errs.AppError)
	// pointer in case we need nil
	ById(string) (*Customer, *errs.AppError)
}
