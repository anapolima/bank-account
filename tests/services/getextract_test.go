package test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/tests/mocks"
)

func TestGetExtractServiceSucceeded(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewExtractService(services.ExtractService{
		GetExtract: mocks.GetExtractSucceeded,
		GetAccount: mocks.GetAccountSucceeded,
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

	want := models.Extract{
		AgencyNumber:            account.AgencyNumber,
		AgencyVerificationCode:  account.AccountVerificationCode,
		AccountNumber:           account.AccountNumber,
		AccountVerificationCode: account.AccountVerificationCode,
		Owner:                   "Test Owner",
		Document:                account.Document,
		Birthdate:               "2012-12-12",
		Balance:                 50.51,
		Transactions: []models.ExtractInfo{
			{
				TransactionID: "transaction-id",
				Type:          "deposit fee",
				Value:         10,
				Date:          "2012-12-12",
			},
			{
				TransactionID: "transaction-id",
				Type:          "deposit",
				Value:         1000,
				Date:          "2012-12-12",
			},
		},
	}

	got, _ := s.GetExtract(account)

	errors := s.GetExtractService(account)
	if errors != nil {
		t.Errorf("expect get extract to succeed, got error %v", errors)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetExtractServiceFailed(t *testing.T) {
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
