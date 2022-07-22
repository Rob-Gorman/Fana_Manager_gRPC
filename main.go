package main

import (
	"fmt"
	"manager/api"
	"manager/configs"
	"manager/publisher"
	"net/http" 
	"os"
)

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	fmt.Printf("\nServing following flag configuration on PORT %s\n", PORT)

	publisher.CreateRedisClient()
	fmt.Printf("\nRedis publisher client connected at %s\n", publisher.Redis.Options().Addr)

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
