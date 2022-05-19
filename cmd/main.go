package main

import (
	"github.com/anapolima/bank-account/app/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
	r := router.Router()
    fmt.Println("Starting server on the port 8080...")

    log.Fatal(http.ListenAndServe(":8080", r))
}