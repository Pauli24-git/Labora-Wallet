package controllers

import (
	"Labora-Wallet/db"
	"Labora-Wallet/models"
	"Labora-Wallet/services"
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

	database, err := db.GetDatabase()
	dbHandlerpost := &db.PostgresDBHandler{Db: database}
	walletService := &services.WalletService{DbHandler: dbHandlerpost}

	err = walletService.ProcessWalletRequest(wallet)

}
