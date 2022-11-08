package api

import (
	"net/http"
)

type CreateTextRequest struct {
	Text string `json:"text"`
}

func writeJsonToResponse(rw *http.ResponseWriter, statusCode int, json []byte) {
	(*rw).Header().Set("Content-Type", "application/json")
	(*rw).WriteHeader(statusCode)
	_, err := (*rw).Write(json)
	if err != nil {
		http.Error(*rw, err.Error(), http.StatusBadRequest)
		return
	}
}
