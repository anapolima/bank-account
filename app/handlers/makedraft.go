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

// MakeDraftHandler handle the request to make a draft
func MakeDraftHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request to make dratf")

	s := services.NewDraftService(services.DraftService{
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

	draft := new(models.Draft)

	log.Printf("Decoding request body")
	err := json.NewDecoder(r.Body).Decode(&draft)

	if err != nil {
		log.Printf("Error decoding request body")
		log.Panic(err)
	}

	s.MakeDraftService(*draft)
}
