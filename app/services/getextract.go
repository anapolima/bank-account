package services

import (
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/models"
)

type ExtractService struct {
	models.Service
	GetExtract models.GetExtract
	GetAccount models.GetAccount
}

func NewExtractService(cfg ExtractService) *ExtractService {
	return &cfg
}

// GetExtractService validate received data to get extract
func (s *ExtractService) GetExtractService(account models.Account) []string {
	log.Printf("Starting to get extract...")
	errorMessages := make([]string, 0)

	account, err := s.GetAccount(account)

	if err != nil {
		log.Printf("An error ocurred while getting account: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)

		return errorMessages
	}

	extract, err := s.GetExtract(account)

	if err != nil {
		log.Printf("An error ocurred while getting extract: %v", err)
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)

		return errorMessages
	}

	log.Printf("Extract fecthed successfully")
	s.Response(s.RWriter, http.StatusOK, extract)

	return nil
}
