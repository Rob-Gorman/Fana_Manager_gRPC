package main

// import (
// 	"context"
// 	"fmt"
// 	"manager/config"
// 	// "github.com/go-redis/redis/v8"
// )
// var ctx = context.TODO()

// type Subscriber struct {

// }

// func (s *Subscriber) subscribeToChannel(channel string) {
// 	pubsub := config.Redis.Subscribe(ctx, channel)
// 	 ch := pubsub.Channel()
// 	 	// // Consume messages
// 	for msg := range ch {
// 		fmt.Printf("%s: %s", msg.Channel, msg.Payload)
// 	}
// }
// func CreateSubscriber() /*return?*/ {

// 	sub := config.Redis.Subscribe(ctx, "flag-toggle-channel")
// 	fmt.Printf("sub created %v\n", sub)
// 	// // Go channel which receives messages
// 	ch := sub.Channel() // what does this documentation mean?

// 	// // Consume messages
// 	for msg := range ch {
// 		fmt.Printf("%s: %s", msg.Channel, msg.Payload)
// 	}
// }
