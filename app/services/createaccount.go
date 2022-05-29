package services

import (
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/models"
)

type CreateAccountService struct {
	models.Service
	NewUUID          models.NewUUID
	CreateAccount    models.CreateAccount
	GenAccount       models.GenAccount
	GenDigit         models.GenDigit
	GenAgency        models.GenAgency
	ValidatePassword models.ValidatePassword
	ValidateDocument models.ValidateDocument
	ValidateDate     models.ValidateDate
}

func NewCreateAccountService(cfg CreateAccountService) *CreateAccountService {
	return &cfg
}

// CreateAccountService validate received data to create a new account
func (s *CreateAccountService) CreateAccountService(userData models.UserData) []string {
	log.Printf("Starting creating account")
	errorMessages := make([]string, 0)
	log.Printf("Start user data check...")

	document, err := s.ValidateDocument(userData.Document)

	if err != nil {
		log.Printf("Invalid document: %v", err)
		errorMessages = append(errorMessages, err.Error())
	}

	birthdate, err := s.ValidateDate(userData.Birthdate)

	if err != nil {
		log.Printf("Invalid birthdate: %v", err)
		errorMessages = append(errorMessages, err.Error())
	}

	password, err := s.ValidatePassword(userData.AccountPassword)

	if err != nil {
		log.Printf("Invalid account password: %v", err)
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		s.ResponseError(s.RWriter, http.StatusBadRequest, errorMessages)
		return errorMessages
	}

	acc := &models.Account{
		AccountID:               s.NewUUID(),
		AgencyNumber:            s.GenAgency(),
		AgencyVerificationCode:  s.GenDigit(),
		AccountNumber:           s.GenAccount(),
		AccountVerificationCode: s.GenDigit(),
		Owner:                   userData.Name,
		Document:                document,
		Birthdate:               birthdate,
		AccountPassword:         password,
	}

	account, err := s.CreateAccount(*acc)

	if err != nil {
		errorMessages = append(errorMessages, err.Error())
		s.ResponseError(s.RWriter, http.StatusInternalServerError, errorMessages)
		return errorMessages
	}

	s.Response(s.RWriter, http.StatusCreated, account)

	return nil
}
