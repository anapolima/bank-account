package handlers

import (
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/app/utils"
	"github.com/gorilla/schema"
)

// GetExtractHandler handle the request for getting account extract
func GetExtractHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting request to get extract")
	errorMessages := make([]string, 0)
	decoder := schema.NewDecoder()
	account := new(models.Account)

	s := services.NewExtractService(services.ExtractService{
		GetAccount: postgresdb.GetAccount,
		GetExtract: postgresdb.GetExtract,
		Service: models.Service{
			RWriter:       w,
			Response:      utils.WriteResponse,
			ResponseError: utils.WriteResponseError,
		},
	})

	log.Printf("Decoding request query params")
	err := decoder.Decode(account, r.URL.Query())

	if err != nil {
		log.Println("Error in GET parameters : ", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

	} else {
		s.GetExtractService(*account)
	}
}
