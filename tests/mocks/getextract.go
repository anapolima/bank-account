package mocks

import (
	"errors"

	"github.com/anapolima/bank-account/app/models"
)

func GetExtractSucceeded(account models.Account) (models.Extract, error) {
	extract := models.Extract{
		AgencyNumber:            account.AgencyNumber,
		AgencyVerificationCode:  account.AccountVerificationCode,
		AccountNumber:           account.AccountNumber,
		AccountVerificationCode: account.AccountVerificationCode,
		Owner:                   "Test Owner",
		Document:                account.Document,
		Birthdate:               "2012-12-12",
		Balance:                 50.51,
		Transactions: []models.ExtractInfo{
			{
				TransactionID: "transaction-id",
				Type:          "deposit fee",
				Value:         10,
				Date:          "2012-12-12",
			},
			{
				TransactionID: "transaction-id",
				Type:          "deposit",
				Value:         1000,
				Date:          "2012-12-12",
			},
		},
	}

	return extract, nil
}

func GetExtractFailed(account models.Account) (models.Extract, error) {
	return (models.Extract{}), errors.New("an error ocurred while getting extract")
}
