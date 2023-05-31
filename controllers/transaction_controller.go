package controllers

import (
	"Labora-Wallet/db"
	"Labora-Wallet/models"
	"Labora-Wallet/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateTransaction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var transactionData models.Transactions
	err := json.NewDecoder(request.Body).Decode(&transactionData)

	if err != nil {
		json.NewEncoder(response).Encode("Hubo un error en la transaccion: " + err.Error())
		return
	}

	database, err := db.GetDatabase()

	TransactionDBHandlerpost := &db.TransactionDBHandler{Db: database}

	transactionService := services.TransactionService{TransactionDBHandler: TransactionDBHandlerpost}

	err = transactionService.CreateTransaction(transactionData)
	if err != nil {
		json.NewEncoder(response).Encode("No se pudo concretar la transaccion: " + err.Error())
		return
	}

	json.NewEncoder(response).Encode("Transacción exitosa.")
	response.WriteHeader(http.StatusOK)
}

func TransactionStatus(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("id")
	idConvertido, err := strconv.Atoi(id)

	database, err := db.GetDatabase()

	TransactionDBHandlerpost := &db.TransactionDBHandler{Db: database}

	transactionService := services.TransactionService{TransactionDBHandler: TransactionDBHandlerpost}

	res, err := transactionService.TransactionsStatusbyID(idConvertido)
	if err != nil {
		json.NewEncoder(response).Encode("Hubo un error en el procesamiento de el Status:" + err.Error())
	}
	json.NewEncoder(response).Encode(res)
}
