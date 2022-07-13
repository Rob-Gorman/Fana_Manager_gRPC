package utils

import (
	"context"
	"log"
	"net/http"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatal(err, msg)
	}
}

func StandardContext() (context.Context, *http.Request) {
	return context.TODO(), nil
	// return context.WithTimeout(context.Background(), 10*time.Second)
}
