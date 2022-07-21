package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"manager/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host string
	db int //index between 0 and 15
	expires time.Duration // expiration time for all elements in cache in seconds
}

func NewRedisCache(host string, db int, exp time.Duration) FlagCache {
	return &redisCache{
		host: host,
		db: db,
		expires: exp,
	}
}
// create a new redis client 
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: "",
		DB: cache.db,
	})
}

// associate flag json to the key
func (cache *redisCache) Set(key string, value *models.Flag) {
	client := cache.getClient()

	// serialize the flag
	json, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Set from redis: marshalling error")
		panic(err)
	}
	// set the key to marshalled data
	client.Set(context.TODO(), key, json, cache.expires*time.Second)
 }

 // get flag based on key
func (cache *redisCache) Get(key string) *models.Flag {
	client := cache.getClient()

// set the key to marshalled data
	val, err := client.Get(context.TODO(), key).Result()
	if err != nil {
		fmt.Println("Get from redis: couldn't get")
		panic(err)
	}

	// unmarshal and store in a flag struct
	flag := models.Flag{}
	err = json.Unmarshal([]byte(val), &flag)
	if err != nil {
		fmt.Println("Get from redis: unmarshalling error")
		panic(err)
	}
	return &flag
 }