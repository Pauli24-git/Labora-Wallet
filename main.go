package main

import (
	"Labora-Wallet/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/wallet", controllers.CreateWallet).Methods(http.MethodPost)
	router.HandleFunc("/wallet", controllers.UpdateWallet).Methods(http.MethodPut)
	router.HandleFunc("/wallet/", controllers.DeleteWallet).Methods(http.MethodDelete)
	router.HandleFunc("/wallet", controllers.WalletStatus).Methods(http.MethodDelete)

	http.ListenAndServe(":8000", router)

	// hace falta hacer un defer con un close?
}
