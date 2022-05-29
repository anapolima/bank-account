package services

import (
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/models"
)

type DepositService struct {
	models.Service
	GetAccount           models.GetAccount
	CreateTransaction    models.CreateTransaction
	CreateFee            models.CreateFee
	UpdateAccountBalance models.UpdateAccountBalance
}

func NewDepositService(cfg DepositService) *DepositService {
	return &cfg
}

// MakeDepositService validate received data to make a deposit
func (s *DepositService) MakeDepositService(deposit models.Deposit) []string {
	log.Printf("Start creating deposit")
	errorMessages := make([]string, 0)

	if deposit.Value <= 0 {
		log.Printf("Cannot deposit null or negative value")
		errorMessages = append(errorMessages, "cannot deposit null or negative value")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)
		return errorMessages
	}

	account, err := postgresdb.GetAccount(deposit.Account)

	if err != nil {
		log.Printf("An error ocurred while getting account: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}

	transaction := &models.Transaction{
		Type:  "deposit",
		Value: deposit.Value,
		Account: models.AccountResponse{
			AgencyNumber:            deposit.Account.AgencyNumber,
			AgencyVerificationCode:  deposit.Account.AgencyVerificationCode,
			AccountNumber:           deposit.Account.AccountNumber,
			AccountVerificationCode: deposit.Account.AccountVerificationCode,
			AccountID:               account.AccountID,
			Document:                deposit.Account.Document,
			Owner:                   account.Owner,
		},
	}

	t, err := s.CreateTransaction(*transaction)

	if err != nil {
		log.Printf("An error ocurred while cerating transaction: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)

		return errorMessages
	}
	_, err = s.UpdateAccountBalance(transaction.Account.AccountID, transaction.Value)

	if err != nil {
		log.Printf("Error updating account balance: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)

		return errorMessages
	}
	_, err = s.CreateFee(*transaction)

	if err != nil {
		log.Printf("Error creating deposit fee")
	}

	t.Account.AccountID = ""
	log.Printf("Transaction created successfully")
	s.Response(s.RWriter, http.StatusCreated, t)

	return nil
}
