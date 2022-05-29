package postgresdb

import (
	"errors"
	"log"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

// GetExtract fetches the account data and its transactions from the database
func GetExtract(account models.Account) (models.Extract, error) {
	log.Printf("Starting getting extract from database")

	db, err := CreateDatabaseConnection()

	if err != nil {
		log.Printf("Error creating database connection: %v", err)
		return (models.Extract{}), errors.New("error creating database connection")
	}
	defer db.Close()

	q := `
		SELECT
			transaction_id
			, transaction_type
			, value
			, date
		FROM public.transactions
		WHERE
			account_id = $1
		ORDER BY date DESC
	`

	rows, err := db.Query(q, account.AccountID)
	if err != nil {
		log.Printf("An error ocurred while getting extract: %v", err)
		return (models.Extract{}), errors.New("an error ocurred while getting extract")
	}

	extract := new(models.Extract)

	extract.AgencyNumber = account.AgencyNumber
	extract.AgencyVerificationCode = account.AgencyVerificationCode
	extract.AccountNumber = account.AccountNumber
	extract.AccountVerificationCode = account.AccountVerificationCode
	extract.Owner = account.Owner
	extract.Document = account.Document
	extract.Birthdate = account.Birthdate
	extract.Balance = account.Balance

	for rows.Next() {
		extractInfo := new(models.ExtractInfo)

		err := rows.Scan(
			&extractInfo.TransactionID,
			&extractInfo.Type,
			&extractInfo.Value,
			&extractInfo.Date,
		)

		if err != nil {
			log.Printf("error scanning transaction")
		}
		extract.Transactions = append(extract.Transactions, *extractInfo)
	}

	return *extract, nil
}
