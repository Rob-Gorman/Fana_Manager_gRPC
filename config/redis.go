package config

import (
	"context"
	"fmt"
	"log"
	"manager/configs"
	"manager/utils"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func CreateRedisClient() {

	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT")),
		Password: configs.GetEnvVar("REDIS_PW"),
		DB:       0, // default
	})
	pong, err := redis.Ping(context.TODO()).Result()

	if err != nil {
		log.Println(pong)
		utils.HandleErr(err, ": Couldn't reach redis server...")
		return
	}
	Redis = redis
}
