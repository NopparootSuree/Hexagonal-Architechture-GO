package service

type NewAccountRequest struct {
	AccountID   string  `json:"account_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountResponse struct {
	AccountID   string  `json:"account_id"`
	AccountType string  `json:"account_type"`
	OpeningDate string  `json:"opening_date"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccontService interface {
	NewAccount(string, NewAccountRequest) (*AccountResponse, error)
	GetAccount(string) ([]AccountResponse, error)
}
