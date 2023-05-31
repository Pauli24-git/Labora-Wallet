package controllers

import (
	"Labora-Wallet/db"
	"Labora-Wallet/models"
	"Labora-Wallet/services"
	"encoding/json"
	"net/http"
	"strconv"
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

	dbHandlerpost := &db.WalletDBHandler{Db: database}
	api := services.API{}
	logs := &db.Logs{Db: database}
	walletService := services.WalletService{WalletDbHandler: dbHandlerpost, Api: api, Logs: logs}

	err = walletService.ProcessWalletRequest(wallet)
	if err != nil {
		json.NewEncoder(response).Encode("Hubo un error en el procesamiento de la Wallet:" + err.Error())
	}

	json.NewEncoder(response).Encode("Se creo con exito.")
}

func UpdateWallet(response http.ResponseWriter, request *http.Request) {
}

func DeleteWallet(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("id")

	database, err := db.GetDatabase()
	idBuscadoConvertido, err := strconv.Atoi(id)

	dbHandlerpost := &db.WalletDBHandler{Db: database}
	api := services.API{}
	logs := db.Logs{Db: database}
	walletService := services.WalletService{WalletDbHandler: dbHandlerpost, Api: api, Logs: &logs}

	err = walletService.ProcessWalletDelete(idBuscadoConvertido)

	if err != nil {
		json.NewEncoder(response).Encode("Hubo un error en la eliminación de la Wallet:" + err.Error())
	}
}

func WalletStatus(response http.ResponseWriter, request *http.Request) {

}
