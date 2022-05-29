package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/app/utils"
	"github.com/google/uuid"
)

// CreateAccountHandler handle the request to create a new account
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request to create account")

	s := services.NewCreateAccountService(services.CreateAccountService{
		NewUUID:          uuid.NewString,
		CreateAccount:    postgresdb.CreateAccount,
		GenAccount:       utils.GenerateAccountNumber,
		GenDigit:         utils.GenerateDigit,
		GenAgency:        utils.GenerateAgencyNumber,
		ValidatePassword: utils.ValidateAccountPassword,
		ValidateDocument: utils.ValidateDocument,
		ValidateDate:     utils.ValidateDate,
		Service: models.Service{
			RWriter:       w,
			Response:      utils.WriteResponse,
			ResponseError: utils.WriteResponseError,
		},
	})

	userData := new(models.UserData)

	log.Printf("Decoding request body")
	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		log.Printf("Error decoding request body")
		log.Panic(err)
	}

	s.CreateAccountService(*userData)
}
