package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anapolima/bank-account/app/models"
)

func WriteResponseError(w http.ResponseWriter, statusCode int, messages []string) {
	log.Printf("Start writing response error")
	var d struct{}
	response := &models.Response{
		Messages: messages,
		Data: d,
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	log.Printf("Start writing response")

	response := &models.Response{
		Data: data,
		Messages: make([]string, 0),
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
