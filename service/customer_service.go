package service

import (
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/repository"
)

type customerService struct {
	cusRepo repository.CustomerRepository
}

func NewCustomerService(cusRepo repository.CustomerRepository) CustomerService {
	return customerService{cusRepo: cusRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.cusRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var custReponse []CustomerResponse

	for _, customer := range customers {
		custRes := CustomerResponse{
			Name:   customer.Name,
			Status: customer.Status,
		}
		custReponse = append(custReponse, custRes)
	}

	return custReponse, nil
}

func (s customerService) GetCustomerByName(name string) (*CustomerResponse, error) {
	customer, err := s.cusRepo.GetOne(name)
	if err != nil {
		return nil, err
	}

	custResponse := CustomerResponse{
		Name:   customer.Name,
		Status: customer.Status,
	}
	return &custResponse, nil
}
