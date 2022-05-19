package handler

import (
	"encoding/json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	res := "Hello world"
	json.NewEncoder(w).Encode(res)
}