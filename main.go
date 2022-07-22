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
	fmt.Printf("\nServing spicy flag configurations on PORT %s\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
