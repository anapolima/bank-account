package mocks

import (
	"errors"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

// UpdateAccountBalance updates the account balance on database
func UpdateAccountBalanceSucceeded(accountId string, value float64) (models.AccountResponse, error) {
	account := models.AccountResponse{
		AgencyNumber:            1111,
		AgencyVerificationCode:  1,
		AccountNumber:           112233,
		AccountVerificationCode: 1,
		Owner:                   "Test Owner",
		Document:                "11111111111",
		Birthdate:               "2012-12-12",
		AccountID:               "account-id",
		AccountPassword:         "hash",
		Balance:                 value,
	}

	return account, nil
}

func UpdateAccountBalanceFailed(accountId string, value float64) (models.AccountResponse, error) {
	return (models.AccountResponse{}), errors.New("unable to update account balance")
}
