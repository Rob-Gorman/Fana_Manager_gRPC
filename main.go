package main

import (
	"fmt"
	"manager/api"
	"manager/cache"
	"manager/configs"
	"manager/publisher"
	"net/http" 
	"os"
)
var (
	flagCache cache.FlagCache = cache.NewRedisCache("localhost:6379", 1, 10)
)
func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	fmt.Printf("\nServing following flag configuration on PORT %s\n", PORT)

	publisher.CreateRedisClient()
	fmt.Printf("\nRedis publisher client connected at %s\n", publisher.Redis.Options().Addr)

	fmt.Printf("\nRedis cache client connected: %v\n", flagCache)



	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
