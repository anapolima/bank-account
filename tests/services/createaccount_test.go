package test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/tests/mocks"
)

func TestCreateAccountServiceSucceeded(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewCreateAccountService(services.CreateAccountService{
		GenAccount:       mocks.GenerateAccountNumber,
		CreateAccount:    mocks.CreateAccountSucceeded,
		GenDigit:         mocks.GenerateDigit,
		GenAgency:        mocks.GenerateAgencyNumber,
		ValidatePassword: mocks.ValidateAccountPasswordSucceeded,
		ValidateDocument: mocks.ValidateDocumentSucceeded,
		ValidateDate:     mocks.ValidateDateSucceeded,
		NewUUID:          mocks.NewUUID,
		Service: models.Service{
			Response:      mocks.WriteResponse,
			ResponseError: mocks.WriteResponseError,
			RWriter:       *w,
		},
	})

	userData := models.UserData{
		Name:            "Test Owner",
		Birthdate:       "2012-12-12",
		AccountPassword: "1111",
		Document:        "11111111111",
	}

	acc := models.Account{
		AgencyNumber:            1111,
		AgencyVerificationCode:  1,
		AccountNumber:           112233,
		AccountVerificationCode: 1,
		Owner:                   userData.Name,
		Document:                userData.Document,
		Birthdate:               userData.Birthdate,
		Balance:                 0,
		AccountID:               "uuid",
		AccountPassword:         "hash1111",
	}

	want := models.AccountResponse{
		AgencyNumber:            1111,
		AgencyVerificationCode:  1,
		AccountNumber:           112233,
		AccountVerificationCode: 1,
		Owner:                   userData.Name,
		Document:                userData.Document,
		Birthdate:               userData.Birthdate,
		Balance:                 0,
		AccountID:               "uuid",
		AccountPassword:         "hash1111",
	}

	errors := s.CreateAccountService(userData)
	if errors != nil {
		t.Errorf("expect create account to succeed, got error %v", errors)
	}

	got, _ := s.CreateAccount(acc)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateAccountServiceFailed(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewExtractService(services.ExtractService{
		GetExtract: mocks.GetExtractFailed,
		GetAccount: mocks.GetAccountFailed,
		Service: models.Service{
			Response:      mocks.WriteResponse,
			ResponseError: mocks.WriteResponseError,
			RWriter:       *w,
		},
	})

	account := models.Account{
		AgencyNumber:            1111,
		AgencyVerificationCode:  1,
		AccountNumber:           112233,
		AccountVerificationCode: 1,
		Document:                "11111111111",
	}

	errors := s.GetExtractService(account)
	if errors == nil {
		t.Errorf("expect get extract to fail, got success")
	}

	_, err := s.GetExtract(account)

	if err == nil {
		t.Errorf("expect get extract to fail, got success")
	}
}
