package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/app/utils"
)

// MakeTransferHandler handle the request to make a transfer
func MakeTransferHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request to make transfer")

	s := services.NewTransferService(services.TransferService{
		GetAccount:           postgresdb.GetAccount,
		CreateTransaction:    postgresdb.CreateTransaction,
		CreateFee:            postgresdb.CreateFee,
		UpdateAccountBalance: postgresdb.UpdateAccountBalance,
		TransactionFee:       postgresdb.TransactionFee,
		Service: models.Service{
			RWriter:       w,
			Response:      utils.WriteResponse,
			ResponseError: utils.WriteResponseError,
		},
	})

	transfer := new(models.Transfer)

	log.Printf("Decoding request body")
	err := json.NewDecoder(r.Body).Decode(&transfer)

	if err != nil {
		log.Printf("Error decoding request body")
		log.Panic(err)
	}

	s.MakeTransferService(*transfer)
}
