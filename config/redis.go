package config

import (
	"fmt"
	"manager/configs"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func CreateRedisClient() {
	// opt, err := redis.ParseURL("redis://localhost:6379/0")
	// if err != nil {
	// 	panic(err)
	// }
	redis:=	redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT")),
				Password: configs.GetEnvVar("REDIS_PW"),
				DB:       0, // default
			})
	Redis = redis
}
