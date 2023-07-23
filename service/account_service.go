package service

import (
	"strings"
	"time"

	"github.com/NopparootSuree/Hexagonal-Architechture-GO/errs"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/logs"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/repository"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccontService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID string, request NewAccountRequest) (*AccountResponse, error) {
	//validate input
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		AccountID:   request.AccountID,
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "online",
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		AccountType: newAcc.AccountType,
		OpeningDate: newAcc.OpeningDate,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status}

	return &response, nil
}

func (s accountService) GetAccount(customID string) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			AccountType: account.AccountType,
			OpeningDate: account.OpeningDate,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}
