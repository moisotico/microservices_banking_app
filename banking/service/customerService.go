package service

import "github.com/moisotico/banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// We want to connect Repository and Service interfaces
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
