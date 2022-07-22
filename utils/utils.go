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
	// I really don't know how to use this. Copy/Pasted from internet til it worked
	return context.TODO(), nil
	// return context.WithTimeout(context.Background(), 10*time.Second)
}
