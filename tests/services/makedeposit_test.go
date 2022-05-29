package test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/anapolima/bank-account/app/models"
	"github.com/anapolima/bank-account/app/services"
	"github.com/anapolima/bank-account/tests/mocks"
)

func TestMakeDepositServiceSucceeded(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewDepositService(services.DepositService{
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

	deposit := models.Deposit{
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
			AgencyNumber:            deposit.Account.AgencyNumber,
			AgencyVerificationCode:  deposit.Account.AccountVerificationCode,
			AccountNumber:           deposit.Account.AccountNumber,
			AccountVerificationCode: deposit.Account.AccountVerificationCode,
			Document:                deposit.Account.Document,
		},
		Value: deposit.Value,
		Type:  "deposit",
	}

	want := models.Transaction{
		ID: "transaction-id",
		Account: models.AccountResponse{
			AgencyNumber:            deposit.Account.AgencyNumber,
			AgencyVerificationCode:  deposit.Account.AccountVerificationCode,
			AccountNumber:           deposit.Account.AccountNumber,
			AccountVerificationCode: deposit.Account.AccountVerificationCode,
			Document:                deposit.Account.Document,
		},
		Value: 100,
		Type:  "deposit",
		Date:  "2012-12-12",
	}

	ts, errors := s.CreateTransaction(transaction)

	if errors != nil {
		t.Errorf("expect make deposit to succeed, got error %v", errors)
	}

	if !reflect.DeepEqual(ts, want) {
		t.Errorf("got %v, want %v", ts, want)
	}
}

func TestMakeDepositServiceFailed(t *testing.T) {
	w := new(http.ResponseWriter)
	s := services.NewDepositService(services.DepositService{
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

	deposit := models.Deposit{
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
			AgencyNumber:            deposit.Account.AgencyNumber,
			AgencyVerificationCode:  deposit.Account.AccountVerificationCode,
			AccountNumber:           deposit.Account.AccountNumber,
			AccountVerificationCode: deposit.Account.AccountVerificationCode,
			Document:                deposit.Account.Document,
		},
		Value: deposit.Value,
		Type:  "deposit",
	}

	_, errors := s.CreateTransaction(transaction)

	if errors == nil {
		t.Errorf("expect make deposit to fail, got success")
	}
}
