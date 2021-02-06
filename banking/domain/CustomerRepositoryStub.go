package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAlll function
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

//helper function
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Moi", City: "San Pedro", Zipcode: "10105", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Woi", City: "San Pedro", Zipcode: "10105", DateOfBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
