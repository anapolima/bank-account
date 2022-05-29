package test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/tests/mocks"
)

func TestMakeDraftServiceSucceeded(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewDraftService(services.DraftService{
		GetAccount:           mocks.GetAccountSucceeded,
		CreateTransaction:    mocks.CreateTransactionSucceeded,
		CreateFee:            mocks.CreateFeeSucceeded,
		UpdateAccountBalance: mocks.UpdateAccountBalanceSucceeded,
		Service: models.Service{
			Response:      mocks.WriteResponse,
			ResponseError: mocks.WriteResponseError,
			RWriter:       *w,
		},
	})

	draft := models.Draft{
		Account: models.Account{
			AgencyNumber:            1111,
			AgencyVerificationCode:  1,
			AccountNumber:           112233,
			AccountVerificationCode: 1,
			Document:                "11111111111",
		},
		Value: 100,
	}

	transaction := models.Transaction{
		Account: models.AccountResponse{
			AgencyNumber:            draft.Account.AgencyNumber,
			AgencyVerificationCode:  draft.Account.AccountVerificationCode,
			AccountNumber:           draft.Account.AccountNumber,
			AccountVerificationCode: draft.Account.AccountVerificationCode,
			Document:                draft.Account.Document,
		},
		Value: draft.Value * -1,
		Type:  "draft",
	}

	want := models.Transaction{
		ID: "transaction-id",
		Account: models.AccountResponse{
			AgencyNumber:            draft.Account.AgencyNumber,
			AgencyVerificationCode:  draft.Account.AccountVerificationCode,
			AccountNumber:           draft.Account.AccountNumber,
			AccountVerificationCode: draft.Account.AccountVerificationCode,
			Document:                draft.Account.Document,
		},
		Value: -100,
		Type:  "draft",
		Date:  "2012-12-12",
	}

	ts, errors := s.CreateTransaction(transaction)

	if errors != nil {
		t.Errorf("expect make draft to succeed, got error %v", errors)
	}

	if !reflect.DeepEqual(ts, want) {
		t.Errorf("got %v, want %v", ts, want)
	}
}

func TestMakeDraftServiceFailed(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewDraftService(services.DraftService{
		GetAccount:           mocks.GetAccountFailed,
		CreateTransaction:    mocks.CreateTransactionFailed,
		CreateFee:            mocks.CreateFeeFailed,
		UpdateAccountBalance: mocks.UpdateAccountBalanceFailed,
		Service: models.Service{
			Response:      mocks.WriteResponse,
			ResponseError: mocks.WriteResponseError,
			RWriter:       *w,
		},
	})

	draft := models.Draft{
		Account: models.Account{
			AgencyNumber:            1111,
			AgencyVerificationCode:  1,
			AccountNumber:           112233,
			AccountVerificationCode: 1,
			Document:                "11111111111",
		},
		Value: 100,
	}

	transaction := models.Transaction{
		Account: models.AccountResponse{
			AgencyNumber:            draft.Account.AgencyNumber,
			AgencyVerificationCode:  draft.Account.AccountVerificationCode,
			AccountNumber:           draft.Account.AccountNumber,
			AccountVerificationCode: draft.Account.AccountVerificationCode,
			Document:                draft.Account.Document,
		},
		Value: draft.Value * -1,
		Type:  "draft",
	}

	_, errors := s.CreateTransaction(transaction)

	if errors == nil {
		t.Errorf("expect make draft to fail, got success")
	}
}
