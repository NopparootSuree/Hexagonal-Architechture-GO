package service

import (
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/errs"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/logs"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/repository"
	"go.mongodb.org/mongo-driver/mongo"
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
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custReponse := []CustomerResponse{}

	for _, customer := range customers {
		custRes := CustomerResponse{
			Name:   customer.Name,
			Status: customer.Status,
		}
		custReponse = append(custReponse, custRes)
	}

	return custReponse, nil
}

func (s customerService) GetCustomerByID(customerID string) (*CustomerResponse, error) {
	customer, err := s.cusRepo.GetOne(customerID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		Name:   customer.Name,
		Status: customer.Status,
	}

	return &custResponse, nil
}
