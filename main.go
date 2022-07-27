package main

import (
	"fmt"
	"manager/api"
	"manager/configs"
	"manager/dev"
	"manager/publisher"
	"net/http"
	"os"
)

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	dev.RefreshSchema(srv.H.DB)
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")

	publisher.CreateRedisClient()
	fmt.Printf("\nRedis publisher client connected at %s\n", publisher.Redis.Options().Addr)

	fmt.Printf("\nServing flag configuration on PORT %s\n", PORT)
	http.Handle("/", http.FileServer(http.Dir("./build")))
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)

	http.ListenAndServe(":3000", nil)
}
