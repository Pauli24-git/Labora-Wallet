package main

import (
	"Labora-Wallet/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/wallet", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/wallet", controllers.UpdateWallet).Methods("PUT")
	router.HandleFunc("/wallet/", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/wallet", controllers.WalletStatus).Methods("GET")

	http.ListenAndServe(":8000", router)

	// hace falta hacer un defer con un close?
}

// falta ObtainScore
// que haga el CreateWallet dependiendo del score
// en mi package db (insert) hacer los logs de los intentos (igual que el wallet)
// para los logs, intentar hacer una interfaz parecido a wallet
// terminar los demas endpoints
// agregar columnas a mi tabla wallet (balance)
