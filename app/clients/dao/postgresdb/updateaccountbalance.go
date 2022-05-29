package postgresdb

import (
	"errors"
	"log"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

// UpdateAccountBalance updates the account balance on database
func UpdateAccountBalance(accountId string, value float64) (models.AccountResponse, error) {
	log.Printf("Starting updating account on database")

	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Printf("Error creating database connection: %v", err)
		return (models.AccountResponse{}), errors.New("error creating database connection")
	}
	defer db.Close()

	q := `
		UPDATE public.bank_account
		SET balance = balance + $1
		WHERE account_id = $2
		RETURNING
		agency_number
		, agency_verification_code
		, account_number
		, account_verification_code
		, balance
		, owner
		, document
		, birthdate
		, account_id
		, account_password
	`

	row := db.QueryRow(q, value, accountId)
	account := new(models.AccountResponse)

	err = row.Scan(&account.AgencyNumber,
		&account.AgencyVerificationCode,
		&account.AccountNumber,
		&account.AccountVerificationCode,
		&account.Balance,
		&account.Owner,
		&account.Document,
		&account.Birthdate,
		&account.AccountID,
		&account.AccountPassword,
	)

	if err != nil {
		log.Print(err)
		return (models.AccountResponse{}), errors.New("unable to update account balance")
	}

	return *account, nil
}
