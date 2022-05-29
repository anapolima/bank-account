package postgresdb

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

var TransactionFee = map[string]float64{
	"draft":    4, // reais
	"transfer": 1, // real
	"deposit":  1, // percent
}

// CreateFee creates a new fee on database
func CreateFee(transaction models.Transaction) (models.Transaction, error) {
	log.Printf("Starting creating %s fee on database", transaction.Type)

	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Printf("Error creating database connection: %v", err)
		return (models.Transaction{}), errors.New("error creating database connection")
	}
	defer db.Close()

	if transaction.Type == "deposit" {
		transaction.Value = math.Abs(transaction.Value*(TransactionFee[transaction.Type]/100)) * -1
	} else {
		transaction.Value = TransactionFee[transaction.Type] * -1
	}
	transaction.Type = fmt.Sprintf("%s fee", transaction.Type)

	t, err := CreateTransaction(transaction)

	if err != nil {
		log.Printf("an error occurred while creating fee transaction %v", err)
		return (models.Transaction{}), errors.New("an error occurred while creating fee transaction")
	}

	_, err = UpdateAccountBalance(t.Account.AccountID, transaction.Value)

	if err != nil {
		log.Printf("Error updating account balance: %v", err)
		return (models.Transaction{}), errors.New("an error occurred while updating account balance")
	}

	return t, nil
}
