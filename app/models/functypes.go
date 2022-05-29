package models

import "net/http"

type CreateAccount func(account Account) (AccountResponse, error)
type CreateFee func(transaction Transaction) (Transaction, error)
type CreateTransaction func(transaction Transaction) (Transaction, error)
type GetAccount func(account Account) (Account, error)
type GetExtract func(account Account) (Extract, error)
type UpdateAccountBalance func(accountId string, value float64) (AccountResponse, error)

type WResponse func(w http.ResponseWriter, statusCode int, data interface{})
type WResponseError func(w http.ResponseWriter, statusCode int, messages []string)

type ValidatePassword func(password string) (string, error)
type ValidateDocument func(document string) (string, error)
type ValidateDate func(date string) (string, error)
type GenDigit func() int
type GenAgency func() int
type GenAccount func() int

type NewUUID func() string

type Service struct {
	RWriter       http.ResponseWriter
	Response      WResponse
	ResponseError WResponseError
}

type TransactionFee map[string]float64
