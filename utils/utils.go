package utils

import (
	"context"
	"log"
	"time"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

func StandardContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	return ctx
}
