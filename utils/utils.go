package utils

import (
	"context"
	"fmt"
)

func HandleErr(err error, msg string) {
	if err != nil {
		fmt.Println(err, msg)
	}
}

func StandardContext() context.Context {
	return context.TODO()
	// return context.WithTimeout(context.Background(), 10*time.Second)
}
