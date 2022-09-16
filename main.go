package main

import (
	"embed"
	"fmt"
	"manager/api"
	"manager/configs"
	"manager/dev"
	"manager/publisher"
	"os"
)

//go:embed build/static/*
var static embed.FS

//go:embed build/index.html
var index []byte

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer(static, index)
	dev.RefreshSchema(srv.H.DB)
	fmt.Println("Connected to postgres!")

	PORT := os.Getenv("PORT")

	publisher.CreateRedisClient()
	go api.Init(&srv.H)

	h2Srv := api.NewH2Server(srv.Router, fmt.Sprintf(":%s", PORT))

	fmt.Printf("\nServing flag configuration on PORT %s\n", PORT)
	h2Srv.ListenAndServe() // HTTP/2 without TLS
}
