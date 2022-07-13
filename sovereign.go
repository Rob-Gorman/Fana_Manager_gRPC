package main

import (
	"fmt"
	"net/http"

	"sovereign/api"
	"sovereign/configs"
)

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	PORT := configs.Port()
	fmt.Printf("\nServing following flag configuration on PORT %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv)
}
