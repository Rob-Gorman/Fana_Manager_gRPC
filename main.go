package main

import (
	"context"
	"fmt"
	"log"
	"manager/api"
	"manager/config"
	"manager/configs"
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
	fmt.Println("Connected to postgres!")
	PORT := os.Getenv("PORT")
	fmt.Printf("\nServing following flag configuration on PORT %s\n", PORT)

	config.CreateRedisClient()
	pong, err := config.Redis.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = config.Redis.Set(ctx, "name", "otto", 0).Err()
	if err != nil {
		log.Fatal("redis not setting value", err)
	}

	val, err := config.Redis.Get(ctx, "name").Result()
	if err != nil {
		log.Fatal("redis not setting value", err)
	}
	fmt.Println("REDIS SETS VALUES:", val)

	fmt.Printf("Redis %v\n", config.Redis) // yay! Redis<localhost:6364 db:0>

	subscribeToChannel(channel) // ok PubSub(flag-toggle-channel)

	// !!!!!!!!!!!!!!! nothing after this gets run  :((((
	publishTo(channel, []byte("ello poppet 1"))

	fmt.Println("last line of main")
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv)
}
