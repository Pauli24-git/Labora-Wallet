package controllers

import (
	"Labora-Wallet/models"
	"encoding/json"
	"net/http"
)

func CreateWallet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var wallet models.Wallet

	err := json.NewDecoder(request.Body).Decode(&wallet)

	if err != nil {
		json.NewEncoder(response).Encode("Hubo un error en la creacion de la Wallet:" + err.Error())
		return
	}

	//{
	// dni: 38424291
	// countryId: AR
	// date: 23-05-23
	//}
}
