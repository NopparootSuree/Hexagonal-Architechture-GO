package service

type CustomerResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomerByName(string) (*CustomerResponse, error)
}
