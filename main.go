package main

import (
	"fmt"

	"sovereign/configs"
	"sovereign/db"
	"sovereign/utils"
)

func main() {
	configs.LoadDotEnv()
	db := db.Init()
	fmt.Println("Connected to postgres!")
	var kindaJSON = map[string]int{"a": 1, "b": 2}
	for i, v := range kindaJSON {
		fmt.Print(i, v)
	}

	fmt.Print(db)
	closeable, err := db.DB()
	utils.HandleErr(err, "can't close won't close")
	closeable.Close()
	// srv := api.NewServer()
	// PORT := configs.Port()
	// fmt.Printf("\nServing following flag configuration on PORT %d\n", PORT)
	// http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv)
}
