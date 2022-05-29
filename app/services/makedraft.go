package services

import (
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/models"
	"golang.org/x/crypto/bcrypt"
)

type DraftService struct {
	models.Service
	GetAccount           models.GetAccount
	CreateTransaction    models.CreateTransaction
	CreateFee            models.CreateFee
	UpdateAccountBalance models.UpdateAccountBalance
	TransactionFee       models.TransactionFee
}

func NewDraftService(cfg DraftService) *DraftService {
	return &cfg
}

// MakeDraftService validate received data to make a draft
func (s *DraftService) MakeDraftService(draft models.Draft) []string {
	log.Printf("Start creating draft")
	errorMessages := make([]string, 0)

	if draft.Value <= 0 {
		log.Printf("Cannot draft null or negative value")
		errorMessages = append(errorMessages, "cannot draft null or negative value")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)
		return errorMessages
	}

	account, err := s.GetAccount(draft.Account)

	if err != nil {
		log.Printf("An error ocurred while getting account: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}
	if draft.Value+s.TransactionFee["draft"] > account.Balance {
		log.Print("Insufficient balance")
		errorMessages = append(errorMessages, "insufficient balance")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(account.AccountPassword),
		[]byte(draft.Account.AccountPassword),
	)

	if err != nil {
		log.Printf("Invalid password")
		errorMessages = append(errorMessages, "invalid password")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)
		return errorMessages
	}

	transaction := &models.Transaction{
		Type:  "draft",
		Value: draft.Value * -1,
		Account: models.AccountResponse{
			AgencyNumber:            draft.Account.AgencyNumber,
			AgencyVerificationCode:  draft.Account.AgencyVerificationCode,
			AccountNumber:           draft.Account.AccountNumber,
			AccountVerificationCode: draft.Account.AccountVerificationCode,
			AccountID:               account.AccountID,
			Document:                draft.Account.Document,
			Owner:                   account.Owner,
		},
	}

	t, err := s.CreateTransaction(*transaction)

	if err != nil {
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
		log.Printf("An error ocurred while cerating transaction: %v", err)
	}
	_, err = postgresdb.UpdateAccountBalance(transaction.Account.AccountID, transaction.Value)

	if err != nil {
		log.Printf("Error updating account balance: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
		return errorMessages
	}
	_, err = s.CreateFee(*transaction)

	if err != nil {
		log.Printf("Error creating draft fee")
		return errorMessages
	}

	t.Value = t.Value * -1
	log.Printf("Draft made successfully")
	s.Response(s.RWriter, http.StatusOK, t)

	return nil
}
