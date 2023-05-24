package models

type Wallet struct {
	ID        int    `json:"id"`
	CountryId string `json: "countryid"`
	DNI       int    `json: "dni"`
	Date      string `json: "date"`
}
