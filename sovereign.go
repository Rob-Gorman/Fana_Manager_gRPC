package main

import (
	"fmt"

	"sovereign/configs"
)

func main() {
	configs.LoadDotEnv()
	configs.DBConnect()
	fmt.Println("Connected to postgres!")
	// srv := api.NewServer()
	// PORT := configs.Port()
	// fmt.Printf("\nServing following flag configuration on PORT %d\n", PORT)
	// http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv)
}
