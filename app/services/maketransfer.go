package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/models"
	"golang.org/x/crypto/bcrypt"
)

type TransferService struct {
	models.Service
	GetAccount           models.GetAccount
	CreateTransaction    models.CreateTransaction
	CreateFee            models.CreateFee
	UpdateAccountBalance models.UpdateAccountBalance
	TransactionFee       models.TransactionFee
}

func NewTransferService(cfg TransferService) *TransferService {
	return &cfg
}

// MakeTransferService validate received data to make a transfer
func (s *TransferService) MakeTransferService(transfer models.Transfer) []string {
	log.Printf("Start creating transfer")
	errorMessages := make([]string, 0)

	if transfer.Value <= 0 {
		log.Printf("Cannot transfer null or negative value")
		errorMessages = append(errorMessages, "cannot transfer null or negative value")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)
		return errorMessages
	}

	originAccount, err := s.GetAccount(transfer.OriginAccount)
	if err != nil {
		log.Printf("An error ocurred while getting origin account: %v", err)
		errorMessages = append(errorMessages, fmt.Sprintf("origin account: %s", err.Error()))
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}

	destinyAccount, err := postgresdb.GetAccount(transfer.DestinyAccount)
	if err != nil {
		log.Printf("An error ocurred while getting destiny account: %v", err)
		errorMessages = append(errorMessages, fmt.Sprintf("destiny account: %s", err.Error()))
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}

	if transfer.Value+s.TransactionFee["transfer"] > originAccount.Balance {
		log.Print("Insufficient balance")
		errorMessages = append(errorMessages, "insufficient balance")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(originAccount.AccountPassword),
		[]byte(transfer.OriginAccount.AccountPassword),
	)

	if err != nil {
		log.Printf("Invalid password")
		errorMessages = append(errorMessages, "invalid password")
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)
		return errorMessages
	}

	originTransaction := &models.Transaction{
		Type:  "transfer",
		Value: transfer.Value * -1,
		Account: models.AccountResponse{
			AgencyNumber:            transfer.OriginAccount.AgencyNumber,
			AgencyVerificationCode:  transfer.OriginAccount.AgencyVerificationCode,
			AccountNumber:           transfer.OriginAccount.AccountNumber,
			AccountVerificationCode: transfer.OriginAccount.AccountVerificationCode,
			AccountID:               originAccount.AccountID,
			Document:                transfer.OriginAccount.Document,
			Owner:                   originAccount.Owner,
		},
	}

	t, err := s.CreateTransaction(*originTransaction)

	if err != nil {
		log.Printf("An error ocurred while cerating origin transaction: %v", err)
		errorMessages = append(errorMessages, fmt.Sprintf("origin account: %s", err.Error()))
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
	}
	_, err = postgresdb.UpdateAccountBalance(originTransaction.Account.AccountID, originTransaction.Value)

	if err != nil {
		errorMessages = append(errorMessages, fmt.Sprintf("origin account: %s", err.Error()))
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
		return errorMessages
	}
	_, err = s.CreateFee(*originTransaction)

	if err != nil {
		log.Printf("Error creating transfer fee")
		return errorMessages
	}

	destinyTransaction := &models.Transaction{
		Type:  "received transfer",
		Value: transfer.Value,
		Account: models.AccountResponse{
			AgencyNumber:            transfer.DestinyAccount.AgencyNumber,
			AgencyVerificationCode:  transfer.DestinyAccount.AgencyVerificationCode,
			AccountNumber:           transfer.DestinyAccount.AccountNumber,
			AccountVerificationCode: transfer.DestinyAccount.AccountVerificationCode,
			AccountID:               destinyAccount.AccountID,
			Document:                transfer.DestinyAccount.Document,
			Owner:                   destinyAccount.Owner,
		},
	}

	_, err = s.CreateTransaction(*destinyTransaction)

	if err != nil {
		log.Printf("An error ocurred while cerating destiny transaction: %v", err)
		errorMessages = append(errorMessages, fmt.Sprintf("destiny account: %s", err.Error()))
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
	}
	_, err = postgresdb.UpdateAccountBalance(destinyTransaction.Account.AccountID, destinyTransaction.Value)

	if err != nil {
		log.Printf("Error updating destiny account balance: %v", err)
		errorMessages = append(errorMessages, fmt.Sprintf("destiny account: %s", err.Error()))
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
		return errorMessages
	}

	t.Value = t.Value * -1

	transferResponse := new(models.TransferTransaction)
	m, _ := json.Marshal(transfer)
	json.Unmarshal(m, &transferResponse)

	transferResponse.ID = t.ID
	transferResponse.Type = t.Type
	transferResponse.Date = t.Date
	transferResponse.Value = t.Value
	log.Printf("Transfer made successfully")
	s.Response(s.RWriter, http.StatusOK, transferResponse)

	return nil
}
