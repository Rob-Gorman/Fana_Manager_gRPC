package main

import (
	"fmt"
	"manager/api"
	"manager/configs"
	"manager/dev"
	"manager/publisher"
	"os"
)

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	dev.RefreshSchema(srv.H.DB)
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")

	publisher.CreateRedisClient()
	api.Init(&srv.H)

	fmt.Printf("\nServing flag configuration on PORT %s\n", PORT)
	// http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
