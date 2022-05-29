package mocks

import (
	"errors"

	"github.com/anapolima/bank-account/app/models"
)

func GetAccountSucceeded(account models.Account) (models.Account, error) {
	acc := models.Account{
		AgencyNumber:            1111,
		AgencyVerificationCode:  1,
		AccountNumber:           112233,
		AccountVerificationCode: 1,
		Owner:                   "Test Owner",
		Document:                "11111111111",
		Birthdate:               "2012-12-12",
		AccountID:               "account-id",
		AccountPassword:         "hash",
		Balance:                 50.51,
	}
	return acc, nil
}

func GetAccountFailed(account models.Account) (models.Account, error) {
	return (models.Account{}), errors.New("unable to get account from database")
}
