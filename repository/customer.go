package repository

type Customer struct {
	CustomerID  string `bson:"customer_id,omitempty"`
	Name        string `bson:"name,omitempty"`
	Adrress     string `bson:"address,omitempty"`
	DateOfBirth string `bson:"date_of_birth,omitempty"`
	Status      string `bson:"status,omitempty"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetOne(string) (*Customer, error)
}
