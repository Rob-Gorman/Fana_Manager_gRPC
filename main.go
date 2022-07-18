package main

import (
	"fmt"
	"net/http"
	"os"

	"manager/api"
	"manager/configs"
)

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	// fmt.Printf("\nServing following flag configuration on PORT %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
