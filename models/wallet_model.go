package models

type Wallet struct {
	ID        int    `json:"id"`
	countryId string `json: "countryid"`
	DNI       int    `json: "dni"`
	date      string `json: "date"`
}
