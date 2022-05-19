package router

import (
	"github.com/anapolima/bank-account/app/handler"
    "github.com/gorilla/mux"
)

func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/user/{id}", handler.HelloWorld).Methods("GET", "OPTIONS")

    return router
}