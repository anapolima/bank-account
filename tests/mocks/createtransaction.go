package mocks

import (
	"errors"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

func CreateTransactionSucceeded(transaction models.Transaction) (models.Transaction, error) {
	t := models.Transaction{
		ID:    "transaction-id",
		Type:  transaction.Type,
		Value: transaction.Value,
		Date:  "2012-12-12",
		Account: models.AccountResponse{
			AgencyNumber:            transaction.Account.AgencyNumber,
			AgencyVerificationCode:  transaction.Account.AgencyVerificationCode,
			AccountNumber:           transaction.Account.AccountNumber,
			AccountVerificationCode: transaction.Account.AccountVerificationCode,
			Document:                transaction.Account.Document,
			AccountID:               transaction.Account.AccountID,
		},
	}

	return t, nil
}

func CreateTransactionFailed(transaction models.Transaction) (models.Transaction, error) {
	return (models.Transaction{}), errors.New("an error occurred while creating transaction")
}
