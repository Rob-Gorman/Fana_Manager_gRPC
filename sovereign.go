package main

import (
	"fmt"

	"sovereign/configs"
)

func main() {
	configs.LoadDotEnv()
	configs.DBConnect()
	fmt.Println("Connected to postgres!")
	var kindaJSON = map[string]int{"a": 1, "b": 2}
	for i, v := range kindaJSON {
		fmt.Print(i, v)
	}
	// srv := api.NewServer()
	// PORT := configs.Port()
	// fmt.Printf("\nServing following flag configuration on PORT %d\n", PORT)
	// http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv)
}
