package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{Name: "jeerawan", Adrress: "49/22", DateOfBirth: "31/03/1998", Status: "stayed"},
		{Name: "tanapong", Adrress: "54/105", DateOfBirth: "25/05/1979", Status: "stayed"},
		{Name: "kongpop", Adrress: "01/55", DateOfBirth: "14/11/1965", Status: "none"},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetOne(name string) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.Name == name {
			return &customer, nil
		}
	}
	return nil, errors.New("Customer not found")
}
