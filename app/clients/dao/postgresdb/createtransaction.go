package postgresdb

import (
	"errors"
	"log"

	"github.com/anapolima/bank-account/app/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// CreateAccount creates a new transaction on database
func CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	log.Printf("Start creating %s on database", transaction.Type)
	db, err := CreateDatabaseConnection()

	if err != nil {
		log.Printf("Error creating database connection: %v", err)
		return (models.Transaction{}), errors.New("error creating database connection")
	}
	defer db.Close()

	transaction.ID = uuid.NewString()

	q := `
	INSERT INTO public.transactions 
		(
			transaction_id
			, transaction_type
			, agency_number
			, agency_verification_code
			, account_number 
			, account_verification_code
			, document
			, account_id
			, value
		)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	)
	RETURNING date
	`

	row := db.QueryRow(
		q,
		transaction.ID,
		transaction.Type,
		transaction.Account.AgencyNumber,
		transaction.Account.AgencyVerificationCode,
		transaction.Account.AccountNumber,
		transaction.Account.AccountVerificationCode,
		transaction.Account.Document,
		transaction.Account.AccountID,
		transaction.Value,
	)

	if row.Err() != nil {
		log.Printf("An error occurred while creating transaction: %v", row.Err())
		return (models.Transaction{}), errors.New("an error occurred while creating transaction")
	}

	row.Scan(&transaction.Date)

	return transaction, nil
}
