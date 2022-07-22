package main
//  import (
// 	// "fmt"
// 	"log"
// 	"context"
// 	"encoding/json"
// 	"manager/config"
//  )
//  var ctx = context.Background()

//  func publishTo(channel string, message []byte)  {
// 	err := config.Redis.Publish(ctx, channel, message).Err()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func encode(message string) []byte {
// 	json, err := json.Marshal(message)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return json
// }
// type Channel struct {
// 	Name string
// }

// type Publisher struct {
// 	// can add more here
// }
// func Init() *Publisher {
// 	configs.LoadDotEnv()
	
// 	p := &Publisher{
// 		Client: redis.NewClient(&redis.Options{
// 			Addr:     fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT")),
// 			Password: "",
// 			DB:       0, // default
// 		}),
// 	}

// 	return p
// }

// var Pub = Init()

// 	p := publisher.Pub
// fmt.Printf("publisher created: %v %v\n", p, ctx)
	// // publish a message
	// p.PublishTo("flag-toggle-channel", "message from go redis")
	// // err = p.Publish(ctx, "flag-toggle-channel", "message from go redis").Err()