package models

type Transaction struct {
	ID      string          `json:"transactionId"`
	Type    string          `json:"type"`
	Value   float64         `json:"value"`
	Date    string          `json:"date,omitempty"`
	Account AccountResponse `json:"account"`
}

type TransferTransaction struct {
	ID             string          `json:"transactionId"`
	Type           string          `json:"type"`
	Value          float64         `json:"value"`
	Date           string          `json:"date,omitempty"`
	OriginAccount  AccountResponse `json:"originAccount"`
	DestinyAccoint AccountResponse `json:"destinyAccount"`
}

type Transfer struct {
	Value          float64 `json:"value"`
	OriginAccount  Account `json:"originAccount"`
	DestinyAccount Account `json:"destinyAccount"`
}

type Deposit struct {
	Account Account `json:"account"`
	Value   float64 `json:"value"`
}

type Draft struct {
	Account Account `json:"account"`
	Value   float64 `json:"value"`
}
