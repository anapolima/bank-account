package postgresdb

import (
	"errors"
	"log"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

// GetAccount fetches an account from database
func GetAccount(account models.Account) (models.Account, error) {
	log.Printf("Starting getting account from database")

	db, err := CreateDatabaseConnection()

	if err != nil {
		log.Printf("Error creating database connection: %v", err)
		return (models.Account{}), errors.New("error creating database connection")
	}
	defer db.Close()

	q := `
		SELECT
			agency_number
			, agency_verification_code
			, account_number
			, account_verification_code
			, owner
			, document
			, birthdate
			, account_id
			, account_password
			, balance
		FROM public.bank_account
		WHERE
			agency_number = $1
			AND
			agency_verification_code = $2
			AND
			account_number = $3
			AND
			account_verification_code = $4
			AND
			document = $5
	`

	row := db.QueryRow(q,
		account.AgencyNumber,
		account.AgencyVerificationCode,
		account.AccountNumber,
		account.AccountVerificationCode,
		account.Document)

	acc := new(models.Account)
	err = row.Scan(
		&acc.AgencyNumber,
		&acc.AgencyVerificationCode,
		&acc.AccountNumber,
		&acc.AccountVerificationCode,
		&acc.Owner,
		&acc.Document,
		&acc.Birthdate,
		&acc.AccountID,
		&acc.AccountPassword,
		&acc.Balance,
	)

	if err != nil {
		log.Print(err.Error())
		return (models.Account{}), errors.New("unable to get account from database")
	}

	log.Printf("Account fethed successfully")
	return *acc, nil
}
