package models

type Account struct {
	AccountID               string  `json:"accountId,omitempty"`
	AgencyNumber            int     `json:"agencyNumber,omitempty"`
	AgencyVerificationCode  int     `json:"agencyVerificationCode,omitempty"`
	AccountNumber           int     `json:"accountNumber,omitempty"`
	AccountVerificationCode int     `json:"accountVerificationCode,omitempty"`
	Owner                   string  `json:"owner,omitempty"`
	Document                string  `json:"document,omitempty"`
	Birthdate               string  `json:"birthdate,omitempty"`
	Balance                 float64 `json:"balance,omitempty"`
	AccountPassword         string  `json:"accountPassword,omitempty"`
}

type AccountResponse struct {
	AccountID               string  `json:"-"`
	AgencyNumber            int     `json:"agencyNumber,omitempty"`
	AgencyVerificationCode  int     `json:"agencyVerificationCode,omitempty"`
	AccountNumber           int     `json:"accountNumber,omitempty"`
	AccountVerificationCode int     `json:"accountVerificationCode,omitempty"`
	Owner                   string  `json:"owner,omitempty"`
	Document                string  `json:"document,omitempty"`
	Birthdate               string  `json:"birthdate,omitempty"`
	Balance                 float64 `json:"balance,omitempty"`
	AccountPassword         string  `json:"-"`
}

type Extract struct {
	AgencyNumber            int           `json:"agencyNumber,omitempty"`
	AgencyVerificationCode  int           `json:"agencyVerificationCode,omitempty"`
	AccountNumber           int           `json:"accountNumber,omitempty"`
	AccountVerificationCode int           `json:"accountVerificationCode,omitempty"`
	Owner                   string        `json:"owner,omitempty"`
	Document                string        `json:"document,omitempty"`
	Birthdate               string        `json:"birthdate,omitempty"`
	Balance                 float64       `json:"balance,omitempty"`
	Transactions            []ExtractInfo `json:"transactions,omitempty"`
}

type ExtractInfo struct {
	TransactionID string  `json:"transactionId"`
	Type          string  `json:"type"`
	Value         float64 `json:"value"`
	Date          string  `json:"date"`
}
