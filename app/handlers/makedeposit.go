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

// MakeDepositHandler handle the request to make a deposit
func MakeDepositHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request to make deposit")

	s := services.NewDepositService(services.DepositService{
		GetAccount:           postgresdb.GetAccount,
		UpdateAccountBalance: postgresdb.UpdateAccountBalance,
		CreateFee:            postgresdb.CreateFee,
		CreateTransaction:    postgresdb.CreateTransaction,
		Service: models.Service{
			RWriter:       w,
			Response:      utils.WriteResponse,
			ResponseError: utils.WriteResponseError,
		},
	})

	deposit := new(models.Deposit)

	log.Printf("Decoding request body")
	err := json.NewDecoder(r.Body).Decode(&deposit)

	if err != nil {
		log.Printf("Error decoding request body")
		log.Panic(err)
	}

	s.MakeDepositService(*deposit)
}
