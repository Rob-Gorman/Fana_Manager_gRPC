package utils

import (
	"encoding/json"
	"net/http"
)

func PayloadResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	// generic function to send an HTTP Response with payload
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func NoRecordResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Record Not Found generic error response
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}

func CreatedResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}

func UpdatedResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}
