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
		http.Error(response, "Hubo un error en la transaccion.", http.StatusBadRequest)
		return
	}

	database, err := db.GetDatabase()

	TransactionDBHandlerpost := &db.TransactionDBHandler{Db: database}

	transactionService := services.TransactionService{TransactionDBHandler: TransactionDBHandlerpost}

	err = transactionService.CreateTransaction(transactionData)
	if err != nil {
		http.Error(response, "No se pudo concretar la transaccion\nHubo un error en el procesamiento de la Wallet.", http.StatusBadRequest)
		return
	}

	json.NewEncoder(response).Encode("Transacción exitosa.")
	json.NewEncoder(response).Encode("Se creo con exito.")
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
		http.Error(response, "Hubo un error en el procesamiento de el Status. ", http.StatusBadRequest)
	}
	json.NewEncoder(response).Encode(res)
}
