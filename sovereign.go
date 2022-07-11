package main

import (
	"fmt"
	"net/http"

	"sovereign/api"
	"sovereign/data"
	"sovereign/utils"
)

func main() {
	utils.LoadDotEnv()
	data.ConnectDB()
	srv := api.NewServer()
	PORT := utils.Port()
	fmt.Printf("\nServing following flag configuration on PORT %d:\n%v\n", PORT, *(srv.Ruleset))
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), srv)
}
