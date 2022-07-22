package main

import (
	"fmt"
	"net/http"
	"os"

	"manager/api"
	"manager/configs"
	"manager/dev"
)

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	dev.RefreshSchema(srv.H.DB)
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	fmt.Printf("\nServing following flag configuration on PORT %s\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
