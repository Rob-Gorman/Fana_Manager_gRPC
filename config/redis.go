package config

import (
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func CreateRedisClient() {
	opt, err := redis.ParseURL("redis://localhost:6364/0")
	if err != nil {
		panic(err)
	}

	redis := redis.NewClient(opt)
	Redis = redis
}


// type Publisher struct {
// 	*redis.Client
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