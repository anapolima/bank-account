package mocks

import (
	"errors"
	"fmt"
	"math"

	"github.com/anapolima/bank-account/app/models"
	_ "github.com/lib/pq"
)

func CreateFeeSucceeded(transaction models.Transaction) (models.Transaction, error) {
	transactionFee := map[string]float64{
		"draft":    4, // reais
		"transfer": 1, // real
		"deposit":  1, // percent
	}

	if transaction.Type == "deposit" {
		transaction.Value = math.Abs(transaction.Value*(transactionFee[transaction.Type]/100)) * -1
	} else {
		transaction.Value = transactionFee[transaction.Type] * -1
	}
	transaction.Type = fmt.Sprintf("%s fee", transaction.Type)
	transaction.Date = "2012-12-12"
	transaction.ID = "transaction-id"

	return transaction, nil
}

func CreateFeeFailed(transaction models.Transaction) (models.Transaction, error) {
	return (models.Transaction{}), errors.New("an error occurred while creating fee transaction")
}
