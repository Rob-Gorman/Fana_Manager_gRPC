package main

import (
	"context"
	"fmt"
	"manager/api"
	"manager/config"
	"manager/configs"
	"manager/dev"
	"manager/publisher"
	"net/http"
	"os"
)

var ctx = context.Background()

const channel = "flag-toggle-channel"

func publishTo(channel string, message []byte) {
	fmt.Println("published a message!")
	err := config.Redis.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println("error with Publish method in main.go", err)
		panic(err)
	}
}

func subscribeToChannel(channel string) {
	pubsub := config.Redis.Subscribe(ctx, channel)
	// ch := pubsub.Channel()

	// fmt.Printf("pubsub subscribed %v\n", pubsub)
	// fmt.Printf("channel: %v\n", ch)
	// Consume messages
	// for msg := range ch {
	// 	fmt.Printf("%s: %s", msg.Channel, msg.Payload)
	// }
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)
	}

}

func main() {
	configs.LoadDotEnv()
	srv := api.NewServer()
	dev.RefreshSchema(srv.H.DB)
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	fmt.Printf("\nServing following flag configuration on PORT %s\n", PORT)

	publisher.CreateRedisClient()
	fmt.Printf("\nRedis publisher client connected at %s\n", publisher.Redis.Options().Addr)

	fmt.Println("last line of main")

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
