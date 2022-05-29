package models

type UserData struct {
	Name            string `json:"name"`
	Document        string `json:"document"`
	Birthdate       string `json:"birthdate"`
	AccountPassword string `json:"accountPassword"`
}
