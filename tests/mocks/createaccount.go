package mocks

import (
	"errors"

	"github.com/anapolima/bank-account/app/models"
)

func CreateAccountSucceeded(account models.Account) (models.AccountResponse, error) {
	acc := &models.AccountResponse{
		AccountID:               NewUUID(),
		AgencyNumber:            GenerateAgencyNumber(),
		AgencyVerificationCode:  GenerateDigit(),
		AccountNumber:           GenerateAccountNumber(),
		AccountVerificationCode: GenerateDigit(),
		Owner:                   "Test Owner",
		Document:                "11111111111",
		Birthdate:               "2012-12-12",
		AccountPassword:         account.AccountPassword,
	}
	return *acc, nil
}

func CreateAccountFailed(acc models.Account) (models.AccountResponse, error) {
	return (models.AccountResponse{}), errors.New("an error occurred while creating account")
}
