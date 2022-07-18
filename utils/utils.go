package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatal(err, msg)
	}
}

func StandardContext() (context.Context, *http.Request) {
	// I really don't know how to use this. Copy/Pasted from internet til it worked
	return context.TODO(), nil
	// return context.WithTimeout(context.Background(), 10*time.Second)
}

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

func ProcessNameToKeyDisplayName(name string) (key, display string) {
	// TODO

	key, display = name, name
	return key, display
}
