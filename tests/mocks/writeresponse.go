package mocks

import (
	"net/http"
)

func WriteResponseError(w http.ResponseWriter, statusCode int, messages []string) {}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {}
