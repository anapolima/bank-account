package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/clients/dao/postgresdb"
	"github.com/anapolima/bank-account/app/router"
)

func main() {
	log.Print("Sarting database setup")
	postgresdb.CreateBankAccountTable()
	postgresdb.CreateTransactionsTable()
	log.Print("Database setup complete")

	r := router.Router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
