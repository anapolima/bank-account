package router

import (
	"github.com/anapolima/bank-account/app/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/create-account", handlers.CreateAccountHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/deposit", handlers.MakeDepositHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/draft", handlers.MakeDraftHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/transfer", handlers.MakeTransferHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/extract", handlers.GetExtractHandler).Methods("GET", "OPTIONS")

	return router
}
