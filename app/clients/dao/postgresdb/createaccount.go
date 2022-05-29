package postgresdb

import (
	"errors"
	"log"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

// CreateAccount creates a new account on database
func CreateAccount(account models.Account) (models.AccountResponse, error) {
	log.Printf("Start creating account on database")
	db, err := CreateDatabaseConnection()

	if err != nil {
		log.Printf("Error creating database connection: %v", err)
		return (models.AccountResponse{}), errors.New("error creating database connection")
	}
	defer db.Close()

	q := `
	INSERT INTO public.bank_account 
		(
			agency_number
			, agency_verification_code
			, account_number
			, account_verification_code
			, owner
			, document
			, birthdate
			, account_id
			, account_password
		)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	)
	`

	row := db.QueryRow(
		q,
		account.AgencyNumber,
		account.AgencyVerificationCode,
		account.AccountNumber,
		account.AccountVerificationCode,
		account.Owner,
		account.Document,
		account.Birthdate,
		account.AccountID,
		account.AccountPassword,
	)

	if row.Err() != nil {
		log.Printf("An error occurred while creating account on database: %v", row.Err())
		return (models.AccountResponse{}), errors.New("an error occurred while creating account")
	}

	return models.AccountResponse{
		AgencyNumber:            account.AgencyNumber,
		AgencyVerificationCode:  account.AgencyVerificationCode,
		AccountNumber:           account.AccountNumber,
		AccountVerificationCode: account.AccountVerificationCode,
		Owner:                   account.Owner,
		Document:                account.Document,
		Birthdate:               account.Birthdate,
		Balance:                 account.Balance,
	}, nil
}
