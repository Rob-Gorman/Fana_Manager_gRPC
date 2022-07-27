package publisher

import (
	"context"
	"fmt"
	"manager/configs"

	"github.com/go-redis/redis/v8"
)

// type RedisHandler struct {
// 	*redis.Client
// }
var Redis *redis.Client

var ctx = context.TODO()

const channel = "flag-toggle-channel"

func CreateRedisClient() {
	configs.LoadDotEnv()

	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT")),
		Password: configs.GetEnvVar("REDIS_PW"),
		DB:       0, // default
	})

	Redis = redis
}
