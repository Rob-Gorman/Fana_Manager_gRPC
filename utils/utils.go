package utils

import (
	"context"
	"log"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatal(err, msg)
	}
}

func StandardContext() context.Context {
	// I really don't know how to use this. Copy/Pasted from internet til it worked
	return context.TODO()
	// return context.WithTimeout(context.Background(), 10*time.Second)
}
