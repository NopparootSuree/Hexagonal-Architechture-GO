package repository

type Account struct {
	AccountID   string  `bson:"account_id,omitempty"`
	CustomerID  string  `bson:"customer_id,omitempty"`
	AccountType string  `bson:"account_type,omitempty"`
	OpeningDate string  `bson:"opening_date"`
	Amount      float64 `bson:"amount,omitempty"`
	Status      string  `bson:"status,omitempty"`
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(string) ([]Account, error)
}
