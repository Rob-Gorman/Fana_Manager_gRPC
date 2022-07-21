package main

import (
	"context"
	"fmt"
	"manager/api"
	"manager/configs"
	"manager/config"
	"net/http"
	"os"
)
var ctx = context.Background()

func publishTo(channel string, message []byte)  {
	fmt.Println("published a message!")
	err := config.Redis.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println("error with Publish method in main.go", err)
		panic(err)
	}
}

func subscribeToChannel(channel string) {
	pubsub := config.Redis.Subscribe(ctx, channel)
	 ch := pubsub.Channel()

	 fmt.Printf("pubsub subscribed %v\n", pubsub)
	 fmt.Printf("channel: %v\n", ch)
	 	// Consume messages
	for msg := range ch {
		fmt.Printf("%s: %s", msg.Channel, msg.Payload)
	}
}

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	fmt.Printf("\nServing following flag configuration on PORT %s\n", PORT)
	
	const channel = "flag-toggle-channel"
	config.CreateRedisClient() 

	fmt.Printf("Redis %v\n", config.Redis) // yay! Redis<localhost:6364 db:0>
	
	subscribeToChannel(channel) // ok PubSub(flag-toggle-channel)
	
	// !!!!!!!!!!!!!!! nothing after this gets run  :((((
	publishTo(channel, []byte("ello poppet 1"))

	fmt.Println("last line of main")
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
